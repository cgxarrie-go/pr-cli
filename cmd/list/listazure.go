package list

import (
	"fmt"
	"log"
	"strings"

	"github.com/cgxarrie-go/prcli/config"
	"github.com/cgxarrie-go/prcli/domain/errors"
	"github.com/cgxarrie-go/prcli/domain/models"
	"github.com/cgxarrie-go/prcli/services/azure"
	"github.com/cgxarrie-go/prcli/services/azure/status"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

const (
	prIDColLength      int = 8
	prTitleColLength   int = 70
	prCreatedColLength int = 25
	prStatusColLength  int = 10
)

// azlsCmd represents the list command
var listAzureCmd = &cobra.Command{
	Use:        "azure",
	Aliases:    []string{"az"},
	SuggestFor: []string{},
	Short:      "List PRs in azure",
	Long:       `List PRs in azure`,
	Example: fmt.Sprintf("list az -s %s\nl az -s %s\nl az -s %s\n"+
		"l az",
		status.Active.Name(), status.Abandoned.Name(),
		status.Closed.Name()),
	Version: "",
	Run: func(cmd *cobra.Command, args []string) {
		st, _ := cmd.Flags().GetString("status")
		if st == "" {
			st = status.Active.Name()
		}
		runListAzureCmd(cmd, st)
	},
}

func runListAzureCmd(cmd *cobra.Command, state string) {

	azCfg, err := loadConfig()
	if err != nil {
		errors.Print(err)
	}
	svc := azure.NewAzureService(azCfg.CompanyName, azCfg.PAT)
	projectRepos := make(map[string][]string)
	for _, project := range azCfg.Projects {
		projectRepos[project.ID] = project.RepositoryIDs
	}

	azStatus, err := status.FromName(state)
	if err != nil {
		errors.Print(err)
		return
	}

	req := azure.GetPRsRequest{ProjectRepos: projectRepos, Status: azStatus}
	prs, err := svc.GetPRs(req)
	if err != nil {
		log.Fatal(err)
	}
	azlsPrint(prs, azCfg.CompanyName)
}

func init() {
	ListCmd.AddCommand(listAzureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listAzureCmd.Flags().StringP("status", "s", status.Active.Name(), "status of PRs to list")

}

func azlsPrint(prs []models.PullRequest, companyName string) {
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
			"pullrequest/%s", companyName, pr.Project.Name,
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
	err = cfg.Load()
	if err != nil {
		return azcfg, err
	}

	return cfg.Azure, nil
}
