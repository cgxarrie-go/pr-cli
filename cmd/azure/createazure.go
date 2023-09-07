package azure

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/services/azure"
)

func RunCreateAzureCmd(cmd *cobra.Command, tgt, ttl string) error {

	azCfg, err := loadConfig()
	if err != nil {
		return err
	}
	svc := azure.NewAzureCreatePullRequestService(azCfg.PAT)

	req := azure.CreatePRRequest{
		Target: tgt,
		Title:  ttl,
	}

	pr, err := svc.Create(req)
	if err != nil {
		return fmt.Errorf("failed to create PR: %w", err)
	}

	lnk := getPRLink(pr.ID, pr.Organization, pr.Project.ID, pr.Repository.ID,
		pr.ID)

	fmt.Printf("PR created with ID: %s\n", lnk)

	return nil
}
