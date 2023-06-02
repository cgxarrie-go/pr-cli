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

func runCreateAzureCmd(cmd *cobra.Command, repo, src, tgt string) error {

	azCfg, err := loadConfig()
	if err != nil {
		return err
	}
	svc := azure.NewAzureService(azCfg.Organization, azCfg.PAT)

	req := azure.CreatePRRequest{
		Repository: repo,
		Source:     src,
		Target:     tgt,
	}
	err = svc.Create(req)
	if err != nil {
		return fmt.Errorf("failed to create PR: %w", err)
	}

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
