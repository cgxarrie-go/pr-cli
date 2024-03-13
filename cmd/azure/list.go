package azure

import (
	"github.com/cgxarrie-go/prq/internal/azure/listprs"
	"github.com/cgxarrie-go/prq/internal/azure/origin"
	"github.com/cgxarrie-go/prq/internal/models"
	"github.com/cgxarrie-go/prq/internal/remote"
)

func RunListCmd(origins remote.Remotes) (
	prs []models.PullRequest, err error) {

	azCfg := loadConfig()
	originSvc := origin.NewService()
	svc := listprs.NewService(azCfg.PAT, originSvc)

	req := listprs.NewRequest(origins)
	prs, err = svc.GetPRs(req)
	return prs, err

}
