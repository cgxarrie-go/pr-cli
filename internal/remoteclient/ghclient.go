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

type githubClient struct {
	base client
	pat  string
}

type ghClientCreateResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"body"`
	URL         string `json:"html_url"`
	Number      int    `json:"number"`
	IsDraft     bool   `json:"draft"`
}

func newGhClient(pat string) ports.RemoteClient {
	return &githubClient{
		base: newClient(),
		pat:  fmt.Sprintf("`:%s", pat),
	}
}

func (r ghClientCreateResponse) ToSvcResponse() ports.CreatePRSvcResponse {
	return ports.CreatePRSvcResponse{
		ID:          fmt.Sprintf("%d", r.Number),
		Title:       r.Title,
		Description: r.Description,
		URL:         r.URL,
		IsDraft:     r.IsDraft,
		Repository:  "",
		Link:        "",
	}
}

func (c *githubClient) Create(req ports.RemoteClientCreateRequest) (
	resp ports.RemoteClientCreateResponse, err error) {

	azReq, err := c.getCreateRequest(req)
	if err != nil {
		return resp, fmt.Errorf("creating http request: %w", err)
	}

	clResp, err := c.base.doCreate(azReq)
	if err != nil {
		return resp, fmt.Errorf("creating PR via client: %w", err)
	}

	ghResp, ok := clResp.(ghClientCreateResponse)
	if !ok {
		return resp, errors.New("casting response to github response")
	}

	resp = ports.RemoteClientCreateResponse{
		ID:          fmt.Sprintf("%d", ghResp.Number),
		Title:       ghResp.Title,
		Description: ghResp.Description,
		URL:         ghResp.URL,
		IsDraft:     ghResp.IsDraft,
	}

	return
}

func (c *githubClient) getCreateRequest(req ports.RemoteClientCreateRequest) (*http.Request, error) {
	b64PAT := base64.RawStdEncoding.EncodeToString([]byte(c.pat))
	bearer := fmt.Sprintf("Basic %s", b64PAT)

	pullRequest := map[string]interface{}{
		"head":        req.Source.Name(),      // Source branch
		"base":        req.Destination.Name(), // Target branch
		"title":       req.Title,              // Title of PR
		"draft":       req.IsDraft,            // Draft PR
		"description": req.Description,        // Description of PR
	}

	body, err := json.Marshal(pullRequest)
	if err != nil {
		return nil, errors.Wrap(err, "marshalling request body")
	}

	ghReq, err := http.NewRequest("POST", req.Remote.CreatePRsURL(),
		bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "creating http request")
	}

	ghReq.Header.Add("Authorization", bearer)
	ghReq.Header.Add("Accept", "application/vnd.github+json")
	ghReq.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	return ghReq, nil
}
