package list

import (
	"fmt"
	"strings"

	"github.com/muesli/termenv"
	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/config"
	"github.com/cgxarrie-go/prq/domain/models"
	"github.com/cgxarrie-go/prq/services/azure"
	"github.com/cgxarrie-go/prq/services/azure/status"
)

const (
	prIDColLength      int = 8
	prTitleColLength   int = 70
	prCreatedColLength int = 25
	prStatusColLength  int = 10
)

func runListAzureCmd(cmd *cobra.Command, state string) error {

	azCfg, err := loadConfig()
	if err != nil {
		return err
	}
	svc := azure.NewAzureReadPullRequestsService(azCfg.Organization, azCfg.PAT)
	projectRepos := make(map[string][]string)
	for _, project := range azCfg.Projects {
		projectRepos[project.ID] = project.RepositoryIDs
	}

	azStatus, err := status.FromName(state)
	if err != nil {
		return err
	}

	req := azure.GetPRsRequest{ProjectRepos: projectRepos, Status: azStatus}
	prs, err := svc.GetPRs(req)
	if err != nil {
		return err
	}
	azlsPrint(prs, azCfg.Organization)

	return nil
}

func azlsPrint(prs []models.PullRequest, organization string) {
	fmt.Printf("Number of PRs : %d \n", len(prs))
	lastProject := ""
	lastRepository := ""

	for i, pr := range prs {
		if i == 0 {
			fmt.Println(azlsPrintableTitle())
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

		url := fmt.Sprintf("https://dev.azure.com/%s/%s/_git/%s/"+
			"pullrequest/%s", organization, pr.Project.Name,
			pr.Repository.Name, pr.ID)

		lnk := termenv.Hyperlink(url, "open")
		created := fmt.Sprintf("%s (%v-%d-%d)",
			strings.Split(pr.CreatedBy, " ")[0],
			pr.Created.Year(), pr.Created.Month(), pr.Created.Day())
		format := getColumnFormat()
		prInfo := fmt.Sprintf(format, pr.ID, pr.ShortenedTitle(70), created,
			pr.Status, lnk)
		fmt.Println(prInfo)
	}
}

func azlsPrintableTitle() string {

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

func loadConfig() (azcfg config.AzureConfig, err error) {
	cfg := config.GetInstance()
	cfg.Load()
	if err != nil {
		return azcfg, err
	}

	return cfg.Azure, nil
}
