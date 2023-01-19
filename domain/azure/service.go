package azure

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cgxarrie/pr-go/domain/errors"
	"github.com/cgxarrie/pr-go/domain/models"
	"github.com/cgxarrie/pr-go/domain/ports"
)

type azureSvc struct {
	conpanyName string
	pat         string
}

// GetPRs implements ports.ProviderService
func (svc azureSvc) GetPRs(req interface{}) (prs []models.PullRequest, err error) {

	getReq, ok := req.(GetPRsRequest)
	if !ok {
		return prs, errors.NewErrInvalidRequestType(getReq, req)
	}

	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/repositories/"+
		"%s/pullrequests?searchCriteria.status=%d&$top=1001&api-version=5.1",
		svc.conpanyName, getReq.ProjectID, getReq.RepositoryID, getReq.Status)

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

func (svc azureSvc) doGet(url string, resp interface{}) (err error) {
	b64PAT := base64.RawStdEncoding.EncodeToString([]byte(svc.pat))
	bearer := fmt.Sprintf("Basic %s", b64PAT)

	azReq, err := http.NewRequest("GET", url, nil)
	azReq.Header.Add("Authorization", bearer)

	client := &http.Client{}
	azResp, err := client.Do(azReq)
	if err != nil {
		return err
	}

	defer azResp.Body.Close()
	return json.NewDecoder(azResp.Body).Decode(resp)

}

// NewAzureService return new instnce of azure service
func NewAzureService(companyName string, pat string) ports.ProviderService {
	return azureSvc{
		conpanyName: companyName,
		pat:         fmt.Sprintf("`:%s", pat),
	}
}
