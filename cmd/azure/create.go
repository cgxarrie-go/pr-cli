package azure

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/domain/azure/createpr"
	"github.com/cgxarrie-go/prq/domain/azure/origin"
)

func RunCreatCmd(cmd *cobra.Command, tgt, ttl string) error {

	azCfg, err := loadConfig()
	if err != nil {
		return err
	}

	originSvc := origin.NewService()
	svc := createpr.NewService(azCfg.PAT, originSvc)

	req := createpr.NewRequest(tgt,ttl)
	pr, err := svc.Run(req)
	if err != nil {
		return fmt.Errorf("failed to create PR: %w", err)
	}

	lnk := getPRLink(pr.ID, pr.Organization, pr.Project.ID, pr.Repository.ID,
		pr.ID)

	fmt.Printf("PR created with ID: %s\n", lnk)

	return nil
}
