package azure

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/domain/azure/listprs"
	"github.com/cgxarrie-go/prq/domain/azure/origin"
	"github.com/cgxarrie-go/prq/domain/azure/status"
	"github.com/cgxarrie-go/prq/domain/models"
	"github.com/cgxarrie-go/prq/utils"
)

const (
	prIDColLength      int = 8
	prTitleColLength   int = 70
	prCreatedColLength int = 25
	prStatusColLength  int = 10
)

func RunListCmd(cmd *cobra.Command, origins utils.Origins, state string) error {

	azCfg, err := loadConfig()
	if err != nil {
		return err
	}
	originSvc := origin.NewService()
	svc := listprs.NewService(azCfg.PAT, originSvc)

	azStatus, err := status.FromName(state)
	if err != nil {
		return err
	}

	req := listprs.NewRequest(origins, azStatus)
	prs, err := svc.GetPRs(req)
	if err != nil {
		return err
	}
	printList(prs)

	return nil
}

func printList(prs []models.PullRequest) {
	fmt.Printf("Number of PRs : %d \n", len(prs))
	lastProject := ""
	lastRepository := ""

	for i, pr := range prs {
		if i == 0 {
			fmt.Println(prinTableTitle())
		}
		if pr.Project.ID != lastProject {
			fmt.Println(pr.Project.Name)
			lastProject = pr.Project.ID
		}
		if pr.Repository.ID != lastRepository {
			if i != 0 {
				fmt.Println()
			}
			fmt.Printf("    %s\n", pr.Repository.Name)
			lastRepository = pr.Repository.ID
		}

		lnk := getPRLink("open", pr.Orgenization, pr.Project.ID, pr.Repository.ID,
			pr.ID)
		created := fmt.Sprintf("%s (%v-%d-%d)",
			strings.Split(pr.CreatedBy, " ")[0],
			pr.Created.Year(), pr.Created.Month(), pr.Created.Day())
		format := getColumnFormat()
		prInfo := fmt.Sprintf(format, pr.ID, pr.ShortenedTitle(70, pr.IsDraft),
			created, pr.Status, lnk)
		fmt.Println(prInfo)
	}
}

func prinTableTitle() string {

	format := "%s\n    %s\n" + getColumnFormat()
	head := fmt.Sprintf(format, "Project", "Repository", "ID", "Title",
		"Created By", "Status", "Link")
	line := strings.Repeat("-", len(head)+5)

	return fmt.Sprintf("%s\n%s", head, line)
}

func getColumnFormat() string {
	return "        %" + fmt.Sprintf("%d", prIDColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prTitleColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prCreatedColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prStatusColLength) + "s " +
		"| %s"
}
