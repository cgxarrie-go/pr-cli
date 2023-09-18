package github

import (
	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/domain/github/listprs"
	"github.com/cgxarrie-go/prq/domain/github/origin"
	"github.com/cgxarrie-go/prq/domain/models"
	"github.com/cgxarrie-go/prq/utils"
)


func RunListCmd(cmd *cobra.Command, origins utils.Origins) (
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
