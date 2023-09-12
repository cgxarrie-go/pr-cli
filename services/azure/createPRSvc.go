package azure

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cgxarrie-go/prq/domain/errors"
	"github.com/cgxarrie-go/prq/domain/models"
	"github.com/cgxarrie-go/prq/domain/ports"
	"github.com/cgxarrie-go/prq/services/azure/branch"
	"github.com/cgxarrie-go/prq/utils"
)

type createPRSvc struct {
	pat string
}

// NewAzureCreatePullRequestService return new instnce of azure service
func NewAzureCreatePullRequestService(pat string) ports.PRCreator {
	return createPRSvc{
		pat: fmt.Sprintf("`:%s", pat),
	}
}

func (svc createPRSvc) url(organization, projectName, repoName string) (
	string, error) {

	return fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/"+
		"repositories/%s/pullRequests?supportsIterations=true&api-version=7.0",
		organization,
		projectName,
		repoName,
	), nil
}

// Create .
func (svc createPRSvc) Create(req any) (pr models.CreatedPullRequest,
	err error) {

	createReq, ok := req.(CreatePRRequest)
	if !ok {
		return pr, errors.NewErrInvalidRequestType(createReq, req)
	}

	src, err := utils.GitCurrentBranchName()
	if err != nil {
		return pr, fmt.Errorf("getting current branch name: %w", err)
	}

	source := branch.NewBranch(src)

	if createReq.Destination == "" {
		createReq.Destination = "master"
	}
	destination := branch.NewBranch(createReq.Destination)

	if createReq.Title == "" {
		createReq.Title =
			fmt.Sprintf("PR from %s to %s", source.Name(), destination.Name())
	}

	organization, projectName, repoName, err := getRepoParams()
	if err != nil {
		return pr, fmt.Errorf("getting repo params: %w", err)
	}

	svcResp := CreatePRResponse{}
	err = svc.doPOST(source, destination, createReq.Title, true,
		organization, projectName, repoName, &svcResp)
	if err != nil {
		return pr, fmt.Errorf("creating PR: %w", err)
	}

	pr = svcResp.ToPullRequest(organization)

	return pr, nil
}

func (svc createPRSvc) doPOST(src, dest branch.Branch, ttl string, draft bool,
	organization, projectName, repoName string,
	resp *CreatePRResponse) (err error) {

	url, err := svc.url(organization, projectName, repoName)
	if err != nil {
		return fmt.Errorf("getting url: %w", err)
	}

	b64PAT := base64.RawStdEncoding.EncodeToString([]byte(svc.pat))
	bearer := fmt.Sprintf("Basic %s", b64PAT)

	pullRequest := map[string]interface{}{
		"sourceRefName": src.FullName(),  // Source branch
		"targetRefName": dest.FullName(), // Target branch
		"title":         ttl,             // Title of PR
		"isDraft":       draft,           // Draft PR
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
	azReq.Header.Add("Content-Type", "application/json")
	azReq.Header.Add("Host", "dev.azure.com")
	azReq.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/117.0")
	azReq.Header.Add("Accept", "application/json;api-version=5.0-preview.1;excludeUrls=true;enumsAsNumbers=true;msDateFormat=true;noArrayWrap=true")
	azReq.Header.Add("Accept-Encoding", "gzip,deflate,br")
	azReq.Header.Add("Referer", fmt.Sprintf("https://dev.azure.com/%s/%s/_git/"+
		"%s/pullrequestcreate?sourceRef=%s&targetRef=%s"+
		"&sourceRepositoryId=%s&targetRepositoryId=%s", organization,
		projectName, repoName, src.Name(), dest.Name(), repoName, repoName))
	azReq.Header.Add("Origin", "https://dev.azure.com")
	azReq.Header.Add("Connection", "keep-alive")
	azReq.Header.Add("Sec-Fetch-Dest", "empty")
	azReq.Header.Add("Sec-Fetch-Mode", "cors")
	azReq.Header.Add("Sec-Fetch-Site", "same-origin")

	client := &http.Client{}
	azResp, err := client.Do(azReq)
	if err != nil {
		return fmt.Errorf("creating PR via client: %w\nurl: %s\n%+v",
			err, url, pullRequest)
	}

	if azResp.StatusCode != http.StatusCreated {
		respBody, err := ioutil.ReadAll(azResp.Body)
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
