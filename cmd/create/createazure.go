package create

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/config"
	"github.com/cgxarrie-go/prq/services/azure"
)

func runCreateAzureCmd(cmd *cobra.Command, tgt, ttl string) error {

	azCfg, err := loadConfig()
	if err != nil {
		return err
	}
	svc := azure.NewAzureCreatePullRequestService(azCfg.PAT)

	req := azure.CreatePRRequest{
		Target: tgt,
		Title:  ttl,
	}

	id, err := svc.Create(req)
	if err != nil {
		return fmt.Errorf("failed to create PR: %w", err)
	}

	fmt.Printf("PR created with ID: %s\n", id)

	return nil
}

func loadConfig() (azcfg config.AzureConfig, err error) {
	cfg := config.GetInstance()
	cfg.Load()
	if err != nil {
		return azcfg, err
	}

	return cfg.Azure, nil
}
