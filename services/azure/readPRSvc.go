package azure

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"

	"github.com/cgxarrie-go/prq/domain/errors"
	"github.com/cgxarrie-go/prq/domain/models"
	"github.com/cgxarrie-go/prq/domain/ports"
)

type readPRSvc struct {
	conpanyName string
	pat         string
}

// NewAzureReadPullRequestsService return new instnce of azure service
func NewAzureReadPullRequestsService(organization string, pat string) ports.PRReader {
	return readPRSvc{
		conpanyName: organization,
		pat:         fmt.Sprintf("`:%s", pat),
	}
}

func (svc readPRSvc) baseUrl(projectID string) string {
	return fmt.Sprintf("https://dev.azure.com"+
		"/%s/%s/_apis/git/repositories", svc.conpanyName, projectID)
}

// GetPRs implements ports.ProviderService
func (svc readPRSvc) GetPRs(req interface{}) (prs []models.PullRequest, err error) {

	getReq, ok := req.(GetPRsRequest)
	if !ok {
		return prs, errors.NewErrInvalidRequestType(getReq, req)
	}

	g := errgroup.Group{}

	for projectID, repositoryIDs := range getReq.ProjectRepos {
		for _, repositoryID := range repositoryIDs {
			projectID, repositoryID := projectID, repositoryID
			g.Go(func() error {
				url := fmt.Sprintf("%s/%s/pullrequests?searchCriteria."+
					"status=%d&api-version=5.1",
					svc.baseUrl(projectID), repositoryID, getReq.Status)

				azPRs, err := svc.getData(url)
				if err == nil {
					prs = append(prs, azPRs...)
				}
				return err
			})
		}
	}

	return prs, g.Wait()
}

func (svc readPRSvc) getData(url string) (prs []models.PullRequest, err error) {
	azPRs := GetPRsResponse{}
	err = svc.doGet(url, &azPRs)
	if err != nil {
		return
	}

	for _, azPR := range azPRs.Value {
		pr := azPR.ToPullRequest()
		prs = append(prs, pr)
	}

	return
}

func (svc readPRSvc) doGet(url string, resp interface{}) (err error) {
	b64PAT := base64.RawStdEncoding.EncodeToString([]byte(svc.pat))
	bearer := fmt.Sprintf("Basic %s", b64PAT)

	azReq, err := http.NewRequest("GET", url, nil)
	azReq.Header.Add("Authorization", bearer)

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
