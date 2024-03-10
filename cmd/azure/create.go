package azure

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/internal/azure/createpr"
	"github.com/cgxarrie-go/prq/internal/azure/origin"
)

func RunCreatCmd(cmd *cobra.Command, tgt, ttl string, isDraft bool) error {

	azCfg := loadConfig()
	originSvc := origin.NewService()
	svc := createpr.NewService(azCfg.PAT, originSvc)

	req := createpr.NewRequest(tgt, ttl, isDraft)
	pr, err := svc.Run(req)
	if err != nil {
		return fmt.Errorf("failed to create PR: %w", err)
	}

	fmt.Printf("PR created with ID: %s (%s)\n", pr.ID, pr.Link)

	return nil
}
