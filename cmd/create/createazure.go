package create

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/config"
	"github.com/cgxarrie-go/prq/services/azure"
)

const (
	prIDColLength      int = 8
	prTitleColLength   int = 70
	prCreatedColLength int = 25
	prStatusColLength  int = 10
)

func runCreateAzureCmd(cmd *cobra.Command, prj, repo, src, tgt string) error {

	azCfg, err := loadConfig()
	if err != nil {
		return err
	}
	svc := azure.NewAzureCreatePullRequestService(azCfg.Organization, azCfg.PAT)

	req := azure.CreatePRRequest{
		Project:    prj,
		Repository: repo,
		Source:     src,
		Target:     tgt,
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
