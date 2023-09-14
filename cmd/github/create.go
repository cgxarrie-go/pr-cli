package github

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/domain/github/createpr"
	"github.com/cgxarrie-go/prq/domain/github/origin"
)

func RunCreatCmd(cmd *cobra.Command, tgt, ttl string) error {

	ghCfg, err := loadConfig()
	if err != nil {
		return err
	}

	originSvc := origin.NewService()
	svc := createpr.NewService(ghCfg.PAT, originSvc)

	req := createpr.NewRequest(tgt,ttl)
	pr, err := svc.Run(req)
	if err != nil {
		return fmt.Errorf("failed to create PR: %w", err)
	}

	fmt.Printf("PR created with ID: %s (%s)\n", pr.ID, pr.Link)

	return nil
}
