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
	ID          int                     `json:"id"`
	URL         string                  `json:"html_url"`
	Number      int                     `json:"number"`
	Title       string                  `json:"title"`
	Body        string                  `json:"body"`
	Status      string                  `json:"sate"`
	MergeStatus string                  `json:"mergeStatus"`
	User        ghClientGetResponseUser `json:"user"`
	IsDraft     bool                    `json:"draft"`
	Created     time.Time               `json:"created_at"`
	Closed      time.Time               `json:"closed_at"`
}

type ghClientGetResponseUser struct {
	Login string `json:"login"`
}

func newGhClient(r ports.Remote, pat string) ports.RemoteClient {
	return &githubClient{
		base: newClient(r),
		pat:  pat,
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

	clResp := ghClientCreateResponse{}
	err = c.base.doCreate(clReq, &clResp)
	if err != nil {
		return resp, errors.Wrap(err, "creating PR in Github")
	}

	resp = ports.RemoteClientCreateResponse{
		ID:          fmt.Sprintf("%d", clResp.Number),
		Title:       clResp.Title,
		Description: clResp.Description,
		URL:         clResp.URL,
		IsDraft:     clResp.IsDraft,
	}

	return
}

func (c *githubClient) Get() (
	resp []ports.RemoteClientGetResponse, err error) {

	clReq, err := c.getGetRequest()
	if err != nil {
		return resp, errors.Wrap(err, "creating http request")
	}

	clResp := []ghClientGetResponse{}
	err = c.base.doGet(clReq, &clResp)
	if err != nil {
		return resp, errors.Wrap(err, "getting PRs from Github")
	}

	resp = make([]ports.RemoteClientGetResponse, len(clResp))
	for i := 0; i < len(clResp); i++ {
		resp[i] = ports.RemoteClientGetResponse{
			ID:          strconv.Itoa(clResp[i].Number),
			Title:       clResp[i].Title,
			Description: clResp[i].Body,
			Status:      clResp[i].Status,
			CreatedBy:   clResp[i].User.Login,
			IsDraft:     clResp[i].IsDraft,
			Created:     clResp[i].Created,
		}
	}

	return
}

func (c *githubClient) getGetRequest() (
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
		"head":  req.Source.Name(),      // Source branch
		"base":  req.Destination.Name(), // Target branch
		"title": req.Title,              // Title of PR
		"draft": req.IsDraft,            // Draft PR
		"body":  req.Description,        // Description of PR
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
