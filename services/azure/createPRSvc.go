package azure

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cgxarrie-go/prq/domain/errors"
	"github.com/cgxarrie-go/prq/domain/ports"
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

func (svc createPRSvc) url() (string, error) {

	companyName, projectName, repoName, err := getRepoParams()
	if err != nil {
		return "", fmt.Errorf("getting repo params: %w", err)
	}

	return fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/"+
		"repositories/%s/pullRequests?supportsIterations=true&api-version=7.0",
		companyName,
		projectName,
		repoName,
	), nil
}

// Create .
func (svc createPRSvc) Create(req interface{}) (id string, err error) {

	createReq, ok := req.(CreatePRRequest)
	if !ok {
		return "", errors.NewErrInvalidRequestType(createReq, req)
	}

	src, err := utils.GitCurrentBranchName()
	if err != nil {
		return "", fmt.Errorf("getting current branch name: %w", err)
	}

	if !strings.HasPrefix(src, "refs/heads/") {
		src = fmt.Sprintf("refs/heads/%s", src)
	}

	if createReq.Target == "" {
		createReq.Target = "refs/heads/master"
	}
	if !strings.HasPrefix(createReq.Target, "refs/heads/") {
		createReq.Target = fmt.Sprintf("refs/heads/%s", createReq.Target)
	}

	if createReq.Title == "" {
		createReq.Title =
			fmt.Sprintf("PR from %s to %s", src, createReq.Target)
	}

	resp := CreatePRResponse{}
	err = svc.doPOST(src, createReq.Target, createReq.Title, true, &resp)
	if err != nil {
		return "", fmt.Errorf("creating PR: %w", err)
	}

	return fmt.Sprintf("%d", resp.ID), nil
}

func (svc createPRSvc) doPOST(src, tgt, ttl string, draft bool,
	resp *CreatePRResponse) (err error) {

	url, err := svc.url()
	if err != nil {
		return fmt.Errorf("getting url: %w", err)
	}

	b64PAT := base64.RawStdEncoding.EncodeToString([]byte(svc.pat))
	bearer := fmt.Sprintf("Basic %s", b64PAT)

	pullRequest := map[string]interface{}{
		"sourceRefName": src, // Source branch
		"targetRefName": tgt, // Target branch
		"title":         ttl, // Title of PR
		"isDraft":       draft,
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
	// azReq.Header.Add("Referer", "https://dev.azure.com/Derivco/Sports-CoreAccount/_git/account/pullrequestcreate?sourceRef=test-branch&targetRef=master&sourceRepositoryId=479b17a4-7d15-48cf-9098-1a7e9dfdaf26&targetRepositoryId=479b17a4-7d15-48cf-9098-1a7e9dfdaf26")
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
		body, err := ioutil.ReadAll(azResp.Body)
		if err != nil {
			body = []byte("cannot read body content")
		}
		return fmt.Errorf("%d - %s", azResp.StatusCode, body)
	}

	defer azResp.Body.Close()
	return json.NewDecoder(azResp.Body).Decode(&resp)

}
