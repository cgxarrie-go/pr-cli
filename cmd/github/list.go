package github

import (
	"github.com/cgxarrie-go/prq/internal/github/listprs"
	"github.com/cgxarrie-go/prq/internal/github/origin"
	"github.com/cgxarrie-go/prq/internal/models"
	"github.com/cgxarrie-go/prq/internal/utils"
)


func RunListCmd(origins utils.Origins) (
	prs []models.PullRequest, err error) {

	ghCfg, err := loadConfig()
	if err != nil {
		return prs, err
	}

	originSvc := origin.NewService()
	svc := listprs.NewService(ghCfg.PAT, originSvc)

	req := listprs.NewRequest(origins)
	prs, err = svc.GetPRs(req)
	return prs, err
}
