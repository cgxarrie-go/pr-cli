package azure

import (
	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/internal/azure/listprs"
	"github.com/cgxarrie-go/prq/internal/azure/origin"
	"github.com/cgxarrie-go/prq/internal/models"
	"github.com/cgxarrie-go/prq/internal/utils"
)

func RunListCmd(cmd *cobra.Command, origins utils.Origins) (
	prs []models.PullRequest, err error) {

	azCfg, err := loadConfig()
	if err != nil {
		return prs, err
	}

	originSvc := origin.NewService()
	svc := listprs.NewService(azCfg.PAT, originSvc)

	req := listprs.NewRequest(origins)
	prs, err = svc.GetPRs(req)
	return prs, err

}

