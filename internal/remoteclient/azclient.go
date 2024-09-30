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
	"github.com/cgxarrie-go/prq/internal/utils"
	"github.com/pkg/errors"
)

type azClient struct {
	base client
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

func newAzClient(r ports.Remote, pat string) ports.RemoteClient {
	return &azClient{
		base: newClient(r),
		pat:  fmt.Sprintf("`:%s", pat),
	}
}

func (c *azClient) Remote() ports.Remote {
	return c.base.remote
}

func (c *azClient) Create(req ports.RemoteClientCreateRequest) (
	resp ports.RemoteClientCreateResponse, err error) {

	azReq, err := c.getCreateRequest(req)
	if err != nil {
		return resp, fmt.Errorf("creating http request: %w", err)
	}

	clResp := azClientCreateResponse{}
	err = c.base.doCreate(azReq, &clResp)
	if err != nil {
		return resp, fmt.Errorf("creating PR via client: %w", err)
	}

	resp = ports.RemoteClientCreateResponse{
		ID:          fmt.Sprintf("%d", clResp.ID),
		Title:       clResp.Title,
		Description: clResp.Description,
		URL:         clResp.URL,
		IsDraft:     clResp.IsDraft,
	}

	return
}

func (c *azClient) Get() (
	resp []ports.RemoteClientGetResponse, err error) {

	clReq, err := c.getGetReques()
	if err != nil {
		return resp, errors.Wrap(err, "creating http request")
	}

	clResp := azClientGetResponse{}
	err = c.base.doGet(clReq, &clResp)
	if err != nil {
		return resp, errors.Wrap(err, "getting PRs from Azure")
	}

	resp = make([]ports.RemoteClientGetResponse, clResp.Count)
	for i := 0; i < clResp.Count; i++ {
		resp[i] = ports.RemoteClientGetResponse{
			ID:          strconv.Itoa(clResp.Value[i].ID),
			Title:       clResp.Value[i].Title,
			Description: clResp.Value[i].Description,
			Status:      clResp.Value[i].Status,
			CreatedBy:   clResp.Value[i].CreatedBy.DisplayName,
			IsDraft:     clResp.Value[i].IsDraft,
			Created:     clResp.Value[i].Created,
		}
	}

	return
}

func (c *azClient) Open(id string) error {
	url := c.base.remote.PRLinkURL(id)
	return utils.OpenBrowser(url)
}

func (c *azClient) getGetReques() (
	*http.Request, error) {

	b64PAT := base64.RawStdEncoding.EncodeToString([]byte(c.pat))
	bearer := fmt.Sprintf("Basic %s", b64PAT)

	clReq, err := http.NewRequest("GET", c.Remote().GetPRsURL(), nil)
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

	clReq, err := http.NewRequest("POST", c.Remote().CreatePRsURL(),
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
