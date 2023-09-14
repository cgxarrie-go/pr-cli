package listprs

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"

	"github.com/cgxarrie-go/prq/domain/azure/origin"
	"github.com/cgxarrie-go/prq/domain/models"
	"github.com/cgxarrie-go/prq/domain/ports"
)

type service struct {
	pat string
	originSvc ports.OriginSvc
}

// NewService return new instnce of azure service
func NewService(pat string, originSvc ports.OriginSvc) ports.PRReader {
	return service{
		pat: fmt.Sprintf("`:%s", pat),
		originSvc: originSvc,
	}
}

// GetPRs implements ports.ProviderService
func (svc service) GetPRs(req ports.ListPRRequest) (
	prs []models.PullRequest, err error) {

	g := errgroup.Group{}

	for _, o := range req.Origins() {
		azOrigin := origin.NewAzureOrigin(o)
		url, err := svc.originSvc.GetPRsURL(o, req.Status())
		if err != nil {
			return prs, fmt.Errorf("gettig url from origin %s: %w", 
			o, err)
		}		
		g.Go(func() error {
			azPRs, err := svc.getData(url, azOrigin.Organizaion())
			if err == nil {
				prs = append(prs, azPRs...)
			}
			return err
		})

	}

	return prs, g.Wait()
}

func (svc service) getData(url, organization string) (
	prs []models.PullRequest, err error) {

	azPRs := Response{}
	err = svc.doGet(url, &azPRs)
	if err != nil {
		return
	}

	for _, azPR := range azPRs.Value {
		pr := azPR.ToPullRequest(organization)
		prs = append(prs, pr)
	}

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
