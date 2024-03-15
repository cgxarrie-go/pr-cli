package remoteclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/pkg/errors"
)

type azClient struct {
	base Client
	pat  string
}

type azClientCreateResponse struct {
	ID          int                              `json:"pullRequestId"`
	Title       string                           `json:"title"`
	Description string                           `json:"description"`
	Repo        azClientCreateResponseRepository `json:"repository"`
	URL         string                           `json:"url"`
	IsDraft     bool                             `json:"isDraft"`
}

type azClientCreateResponseRepository struct {
	ID      string                        `json:"id"`
	Name    string                        `json:"name"`
	URL     string                        `json:"url"`
	Project azClientCreateResponseProject `json:"project"`
}

type azClientCreateResponseProject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type azClientGetResponse struct {
	Value []azClientGetResponseItem `json:"value"`
	Count int                       `json:"count"`
}

type azClientGetResponseItem struct {
	ID          int                         `json:"pullRequestId"`
	Title       string                      `json:"title"`
	Description string                      `json:"description"`
	Status      string                      `json:"status"`
	MergeStatus string                      `json:"mergeStatus"`
	CreatedBy   azClientGetResponseItemUser `json:"createdBy"`
	URL         string                      `json:"url"`
	IsDraft     bool                        `json:"isDraft"`
	Created     time.Time                   `json:"creationDate"`
}
type azClientGetResponseItemUser struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"uniqueName"`
}

func newAzClient(pat string) ports.RemoteClient {
	return &azClient{
		base: newClient(),
		pat:  fmt.Sprintf("`:%s", pat),
	}
}

func (c *azClient) Create(req ports.RemoteClientCreateRequest) (
	resp ports.RemoteClientCreateResponse, err error) {

	azReq, err := c.getCreateRequest(req)
	if err != nil {
		return resp, fmt.Errorf("creating http request: %w", err)
	}

	clResp, err := c.base.doCreate(azReq)
	if err != nil {
		return resp, fmt.Errorf("creating PR via client: %w", err)
	}

	azResp, ok := clResp.(azClientCreateResponse)
	if !ok {
		return resp, errors.New("casting response to github response")
	}

	resp = ports.RemoteClientCreateResponse{
		ID:          fmt.Sprintf("%d", azResp.ID),
		Title:       azResp.Title,
		Description: azResp.Description,
		URL:         azResp.URL,
		IsDraft:     azResp.IsDraft,
	}

	return
}

func (c *azClient) Get(req ports.RemoteClientGetRequest) (
	resp []ports.RemoteClientGetResponse, err error) {

	clReq, err := c.getGetReques(req)
	if err != nil {
		return resp, errors.Wrap(err, "creating http request")
	}

	clResp, err := c.base.doGet(clReq)
	if err != nil {
		return resp, errors.Wrap(err, "getting PRs from Github")
	}

	azResp, ok := clResp.(azClientGetResponse)
	if !ok {
		return resp, errors.New("casting response to github response")
	}

	resp = make([]ports.RemoteClientGetResponse, azResp.Count)
	for i := 0; i < azResp.Count; i++ {
		resp[i] = ports.RemoteClientGetResponse{
			ID:          strconv.Itoa(azResp.Value[i].ID),
			Title:       azResp.Value[i].Title,
			Description: azResp.Value[i].Description,
			Status:      azResp.Value[i].Status,
			CreatedBy:   azResp.Value[i].CreatedBy.DisplayName,
			IsDraft:     azResp.Value[i].IsDraft,
			Created:     azResp.Value[i].Created,
		}
	}

	return
}

func (c *azClient) getGetReques(req ports.RemoteClientGetRequest) (
	*http.Request, error) {

	b64PAT := base64.RawStdEncoding.EncodeToString([]byte(c.pat))
	bearer := fmt.Sprintf("Basic %s", b64PAT)

	clReq, err := http.NewRequest("GET", req.Remote.GetPRsURL(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating http request")
	}
	clReq.Header.Add("Authorization", bearer)

	return clReq, nil
}

func (c *azClient) getCreateRequest(req ports.RemoteClientCreateRequest) (*http.Request, error) {
	b64PAT := base64.RawStdEncoding.EncodeToString([]byte(c.pat))
	bearer := fmt.Sprintf("Basic %s", b64PAT)

	pullRequest := map[string]interface{}{
		"sourceRefName": req.Source.FullName(),      // Source branch
		"targetRefName": req.Destination.FullName(), // Target branch
		"title":         req.Title,                  // Title of PR
		"isDraft":       req.IsDraft,                // Draft PR
		"description":   req.Description,            // Description of PR
	}

	body, err := json.Marshal(pullRequest)
	if err != nil {
		return nil, errors.Wrap(err, "marshalling request body")
	}

	clReq, err := http.NewRequest("POST", req.Remote.CreatePRsURL(),
		bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "creating http request")
	}

	clReq.Header.Add("Authorization", bearer)
	clReq.Header.Add("Content-Type", "application/json")
	clReq.Header.Add("Host", "dev.azure.com")
	clReq.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/117.0")
	clReq.Header.Add("Accept", "application/json;api-version=5.0-preview.1;excludeUrls=true;enumsAsNumbers=true;msDateFormat=true;noArrayWrap=true")
	clReq.Header.Add("Accept-Encoding", "gzip,deflate,br")
	clReq.Header.Add("Referer", "prq")
	clReq.Header.Add("Origin", "https://dev.azure.com")
	clReq.Header.Add("Connection", "keep-alive")
	clReq.Header.Add("Sec-Fetch-Dest", "empty")
	clReq.Header.Add("Sec-Fetch-Mode", "cors")
	clReq.Header.Add("Sec-Fetch-Site", "same-origin")

	return clReq, nil
}
