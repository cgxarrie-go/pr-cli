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

type ghClientGetResponse struct {
	Value []ghClientGetResponseItem `json:"value"`
	Count int                       `json:"count"`
}
type ghClientGetResponseItem struct {
	ID          int                         `json:"id"`
	URL         string                      `json:"html_url"`
	Number      int                         `json:"number"`
	Title       string                      `json:"title"`
	Body        string                      `json:"body"`
	Status      string                      `json:"sate"`
	MergeStatus string                      `json:"mergeStatus"`
	User        ghClientGetResponseItemUser `json:"user"`
	IsDraft     bool                        `json:"draft"`
	Created     time.Time                   `json:"created_at"`
	Closed      time.Time                   `json:"closed_at"`
}

type ghClientGetResponseItemUser struct {
	Login string `json:"login"`
}

func newGhClient(r ports.Remote, pat string) ports.RemoteClient {
	return &githubClient{
		base: newClient(r),
		pat:  fmt.Sprintf("`:%s", pat),
	}
}

func (c *githubClient) Remote() ports.Remote {
	return c.base.remote
}

func (c *githubClient) Create(req ports.RemoteClientCreateRequest) (
	resp ports.RemoteClientCreateResponse, err error) {

	clReq, err := c.getCreateRequest(req)
	if err != nil {
		return resp, errors.Wrap(err, "creating http request")
	}

	clResp, err := c.base.doCreate(clReq)
	if err != nil {
		return resp, errors.Wrap(err, "creating PR in Github")
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

func (c *githubClient) Get() (
	resp []ports.RemoteClientGetResponse, err error) {

	clReq, err := c.getGetReques()
	if err != nil {
		return resp, errors.Wrap(err, "creating http request")
	}

	clResp, err := c.base.doGet(clReq)
	if err != nil {
		return resp, errors.Wrap(err, "getting PRs from Github")
	}

	ghResp, ok := clResp.(ghClientGetResponse)
	if !ok {
		return resp, errors.New("casting response to github response")
	}

	resp = make([]ports.RemoteClientGetResponse, ghResp.Count)
	for i := 0; i < ghResp.Count; i++ {
		resp[i] = ports.RemoteClientGetResponse{
			ID:          strconv.Itoa(ghResp.Value[i].Number),
			Title:       ghResp.Value[i].Title,
			Description: ghResp.Value[i].Body,
			Status:      ghResp.Value[i].Status,
			CreatedBy:   ghResp.Value[i].User.Login,
			IsDraft:     ghResp.Value[i].IsDraft,
			Created:     ghResp.Value[i].Created,
		}
	}

	return
}

func (c *githubClient) getGetReques() (
	*http.Request, error) {

	bearer := fmt.Sprintf("Bearer %s", c.pat)

	clReq, err := http.NewRequest("GET", c.Remote().GetPRsURL(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating http request")
	}
	clReq.Header.Add("Authorization", bearer)
	clReq.Header.Add("Accept", "application/vnd.github+json")
	clReq.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	return clReq, nil
}

func (c *githubClient) getCreateRequest(req ports.RemoteClientCreateRequest) (
	*http.Request, error) {
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

	clReq, err := http.NewRequest("POST", c.Remote().CreatePRsURL(),
		bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "creating http request")
	}

	clReq.Header.Add("Authorization", bearer)
	clReq.Header.Add("Accept", "application/vnd.github+json")
	clReq.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	return clReq, nil
}
