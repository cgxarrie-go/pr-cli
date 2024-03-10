package createpr

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/cgxarrie-go/prq/internal/config"
	"github.com/cgxarrie-go/prq/internal/github/branch"
	"github.com/cgxarrie-go/prq/internal/github/origin"
	"github.com/cgxarrie-go/prq/internal/models"
	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/utils"
)

type service struct {
	pat       string
	originSvc ports.OriginSvc
}

// NewService return new instnce of github service
func NewService(pat string, originSvc ports.OriginSvc) ports.PRCreator {
	return service{
		pat:       fmt.Sprintf("`:%s", pat),
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
	destination :=
		branch.NewBranch(config.GetInstance().Github.DefaultTargetBranch)
	if req.Destination() != "" {
		destination = branch.NewBranch(req.Destination())
	}

	title := req.Title()
	if req.Title() == "" {
		title = fmt.Sprintf("PR from %s to %s",
			source.Name(), destination.Name())
	}

	desc, err := os.ReadFile("./docs/pull_request_template.md")
	if err != nil {
		desc = []byte("")
	}

	o, err := utils.CurrentFolderRemote()
	if err != nil {
		return pr, fmt.Errorf("getting repository origin: %w", err)
	}

	ghOrigin := origin.NewGithubOrigin(o)

	svcResp := Response{}
	err = svc.doPOST(source, destination, title, string(desc), req.IsDraft(),
		ghOrigin, &svcResp)
	if err != nil {
		return pr, fmt.Errorf("creating PR: %w", err)
	}

	pr = svcResp.ToPullRequest(ghOrigin.User())
	pr.Link, _ = svc.originSvc.PRLink(o, pr.ID, "open")

	return pr, nil
}

func (svc service) doPOST(src, dest branch.Branch, ttl, desc string, draft bool,
	o origin.GithubOrigin, resp *Response) (err error) {

	url, err := svc.originSvc.CreatePRsURL(o.Remote)
	if err != nil {
		return fmt.Errorf("getting url: %w", err)
	}

	b64PAT := base64.RawStdEncoding.EncodeToString([]byte(svc.pat))
	bearer := fmt.Sprintf("Basic %s", b64PAT)

	pullRequest := map[string]interface{}{
		"head":        src.Name(),  // Source branch
		"base":        dest.Name(), // Target branch
		"title":       ttl,         // Title of PR
		"draft":       draft,       // Draft PR
		"description": desc,        // Description of PR
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
