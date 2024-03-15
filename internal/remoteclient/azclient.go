package remoteclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cgxarrie-go/prq/internal/ports"
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

// azClientCreateResponseProject pull request response project
type azClientCreateResponseProject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (r azClientCreateResponse) ToSvcResponse() ports.CreatePRSvcResponse {
	return ports.CreatePRSvcResponse{
		ID:          fmt.Sprintf("%d", r.ID),
		Title:       r.Title,
		Description: r.Description,
		URL:         r.URL,
		IsDraft:     r.IsDraft,
		Link:        "",
		Repository:  "",
	}
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

	azReq, err := http.NewRequest("POST", req.Remote.CreatePRsURL(),
		bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "creating http request")
	}

	azReq.Header.Add("Authorization", bearer)
	azReq.Header.Add("Content-Type", "application/json")
	azReq.Header.Add("Host", "dev.azure.com")
	azReq.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/117.0")
	azReq.Header.Add("Accept", "application/json;api-version=5.0-preview.1;excludeUrls=true;enumsAsNumbers=true;msDateFormat=true;noArrayWrap=true")
	azReq.Header.Add("Accept-Encoding", "gzip,deflate,br")
	azReq.Header.Add("Referer", "prq")
	azReq.Header.Add("Origin", "https://dev.azure.com")
	azReq.Header.Add("Connection", "keep-alive")
	azReq.Header.Add("Sec-Fetch-Dest", "empty")
	azReq.Header.Add("Sec-Fetch-Mode", "cors")
	azReq.Header.Add("Sec-Fetch-Site", "same-origin")

	return azReq, nil
}
