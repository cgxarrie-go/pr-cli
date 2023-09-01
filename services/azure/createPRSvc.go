package azure

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cgxarrie-go/prq/domain/errors"
	"github.com/cgxarrie-go/prq/domain/ports"
	"github.com/cgxarrie-go/prq/utils"
)

type createPRSvc struct {
	conpanyName string
	pat         string
}

// NewAzureCreatePullRequestService return new instnce of azure service
func NewAzureCreatePullRequestService(organization string, pat string) ports.PRCreator {
	return createPRSvc{
		conpanyName: organization,
		pat:         fmt.Sprintf("`:%s", pat),
	}
}

func (svc createPRSvc) baseUrl(projectID, repository string) string {
	return fmt.Sprintf(
		"https://dev.azure.com/%s/%s/_apis/git/repositories/%s/pullrequests?"+
			"api-version=7.0&supportsIterations=true",
		svc.conpanyName,
		projectID,
		repository,
	)
}

// Create .
func (svc createPRSvc) Create(req interface{}) (id string, err error) {

	createReq, ok := req.(CreatePRRequest)
	if !ok {
		return "", errors.NewErrInvalidRequestType(createReq, req)
	}

	resp := CreatePRResponse{}

	err = svc.doPOST(createReq, resp)

	if err != nil {
		return "", fmt.Errorf("creating PR: %w", err)
	}

	return fmt.Sprintf("%d", resp.ID), nil
}

func (svc createPRSvc) doPOST(req CreatePRRequest,
	resp interface{}) (err error) {

	url := svc.baseUrl(req.Project, req.Repository)

	b64PAT := base64.RawStdEncoding.EncodeToString([]byte(svc.pat))
	bearer := fmt.Sprintf("Basic %s", b64PAT)

	sourceBranch, err := utils.GitCurrentBranchName()
	if err != nil {
		return fmt.Errorf("getting current branch name: %w", err)
	}

	pullRequest := map[string]interface{}{
		"sourceRefName": sourceBranch, // Source branch
		"targetRefName": req.Target,   // Target branch
		"title":         req.Title,
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

	client := &http.Client{}
	azResp, err := client.Do(azReq)
	if err != nil {
		return err
	}

	if azResp.StatusCode != http.StatusOK {
		return fmt.Errorf("%d - %s", azResp.StatusCode, azResp.Status)
	}

	defer azResp.Body.Close()
	return json.NewDecoder(azResp.Body).Decode(resp)

}
