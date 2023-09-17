package createpr

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cgxarrie-go/prq/domain/github/branch"
	"github.com/cgxarrie-go/prq/domain/github/origin"
	"github.com/cgxarrie-go/prq/domain/models"
	"github.com/cgxarrie-go/prq/domain/ports"
	"github.com/cgxarrie-go/prq/utils"
)

type service struct {
	pat string
	originSvc ports.OriginSvc
}

// NewService return new instnce of github service
func NewService(pat string, originSvc ports.OriginSvc) ports.PRCreator {
	return service{
		pat: fmt.Sprintf("`:%s", pat),
		originSvc: originSvc,
	}
}

// Create .
func (svc service) Run(req ports.CreatePRRequest) (
	pr models.CreatedPullRequest, err error) {

	src, err := utils.GitCurrentBranchName()
	if err != nil {
		return pr, fmt.Errorf("getting current branch name: %w", err)
	}

	source := branch.NewBranch(src)
	destination := branch.Branch{}
	if req.Destination() == "" {
		destination = branch.NewBranch("main")
	} else {
		destination = branch.NewBranch(req.Destination())
	}
	
	title := req.Title()
	if req.Title() == "" {
		title = fmt.Sprintf("PR from %s to %s", 
		source.Name(), destination.Name())
	}

	o, err := utils.CurrentOrigin()
	if err != nil {
		return pr, fmt.Errorf("getting repository origin: %w", err)
	}

	ghOrigin := origin.NewGithubOrigin(o)

	svcResp := Response{}
	err = svc.doPOST(source, destination, title, true, ghOrigin, &svcResp)
	if err != nil {
		return pr, fmt.Errorf("creating PR: %w", err)
	}

	pr = svcResp.ToPullRequest(ghOrigin.User())
	pr.Link, err = svc.originSvc.PRLink(o, pr.ID, "open")
	
	return pr, nil
}

func (svc service) doPOST(src, dest branch.Branch, ttl string, draft bool,
	o origin.GithubOrigin, resp *Response) (err error) {

	url, err := svc.originSvc.CreatePRsURL(o.Origin)
	if err != nil {
		return fmt.Errorf("getting url: %w", err)
	}

	b64PAT := base64.RawStdEncoding.EncodeToString([]byte(svc.pat))
	bearer := fmt.Sprintf("Basic %s", b64PAT)

	pullRequest := map[string]interface{}{
		"head": src.Name(),  // Source branch
		"base": dest.Name(), // Target branch
		"title":         ttl,             // Title of PR
		"draft":       draft,           // Draft PR
	}

	body, err := json.Marshal(pullRequest)
	if err != nil {
		return fmt.Errorf("marshalling request body: %w", err)
	}

	azReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("creating http request: %w", err)
	}

	azReq.Header.Add("Authorization", bearer)
	azReq.Header.Add("Accept", "application/vnd.github+json")
	azReq.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	azResp, err := client.Do(azReq)
	if err != nil {
		return fmt.Errorf("creating PR via client: %w\nurl: %s\n%+v",
			err, url, pullRequest)
	}

	if azResp.StatusCode != http.StatusCreated {
		respBody, err := io.ReadAll(azResp.Body)
		if err != nil {
			respBody = []byte("cannot read response body content")
		}

		return fmt.Errorf("response code: %d\n"+
			"response body: %+v\n"+
			"pull request: %+v\n"+
			"url: %s\n"+
			"request: %+v",
			azResp.StatusCode, string(respBody), pullRequest, url, azReq)

	}

	defer azResp.Body.Close()
	return json.NewDecoder(azResp.Body).Decode(&resp)

}
