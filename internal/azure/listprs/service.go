package listprs

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cgxarrie-go/prq/internal/azure/origin"
	"github.com/cgxarrie-go/prq/internal/models"
	"github.com/cgxarrie-go/prq/internal/ports"
)

type service struct {
	pat       string
	originSvc ports.OriginSvc
}

// NewService return new instnce of azure service
func NewService(pat string, originSvc ports.OriginSvc) ports.PRReader {
	return service{
		pat:       fmt.Sprintf("`:%s", pat),
		originSvc: originSvc,
	}
}

// GetPRs implements ports.ProviderService
func (svc service) GetPRs(req ports.ListPRRequest) (
	prs []models.PullRequest, err error) {

	for _, o := range req.Origins() {
		azOrigin := origin.NewAzureOrigin(o)
		url, err := svc.originSvc.GetPRsURL(o)
		if err != nil {
			return prs, fmt.Errorf("gettig url from origin %s: %w",
				o, err)
		}
		azPRs, err := svc.getData(url)
		if err != nil {
			return prs, fmt.Errorf("getting PRs from %s: %w",
				o, err)
		}

		for _, azPR := range azPRs.Value {
			pr := azPR.ToPullRequest(azOrigin.Organization())
			pr.Link, err = svc.originSvc.PRLink(azOrigin.Origin, pr.ID,
				"open")
			if err != nil {
				return prs, fmt.Errorf("getting PR link from %s: %w",
					o, err)
			}
			prs = append(prs, pr)
		}
	}

	return prs, err
}

func (svc service) getData(url string) (
	prs Response, err error) {

	err = svc.doGet(url, &prs)
	return
}

func (svc service) doGet(url string, resp interface{}) (err error) {
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
