package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/muesli/termenv"
	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prcli/config"
	"github.com/cgxarrie-go/prcli/domain/models"
	"github.com/cgxarrie-go/prcli/services/azure"
)

const (
	prIDColLength      int = 8
	prTitleColLength   int = 70
	prCreatedColLength int = 25
	prStatusColLength  int = 10
)

// azlsCmd represents the list command
var azlsCmd = &cobra.Command{
	Use:        "azure",
	Aliases:    []string{"az"},
	SuggestFor: []string{},
	Short:      "List PRs in azure",
	Long:       `List PRs in azure`,
	Example:    "az list active",
	ValidArgs:  []string{"", "ac", "active", "ab", "abandoned", "cl", "closed"},
	Version:    "",
	Run: func(cmd *cobra.Command, args []string) {
		runAzlsCmd(cmd, args)
	},
}

func runAzlsCmd(cmd *cobra.Command, args []string) {
	if len(args) > 1 {
		fmt.Println("Error : expected one argument")
		return
	}

	if len(args) == 0 || args[0] == "" {
		args = []string{"ac"}
	}

	status := 1
	switch args[0] {
	case "ac", "active":
		status = 1
	case "ab", "abandoned":
		status = 2
	case "cl", "closed":
		status = 3
	default:
		fmt.Printf("Error Invalid argument: %s", args[0])
		return
	}
	cfg := config.GetInstance().Azure
	svc := azure.NewAzureService(cfg.CompanyName, cfg.PAT)
	projectRepos := make(map[string][]string)
	for _, project := range cfg.Projects {
		projectRepos[project.ID] = project.RepositoryIDs
	}
	req := azure.GetPRsRequest{ProjectRepos: projectRepos, Status: status}
	prs, err := svc.GetPRs(req)
	if err != nil {
		log.Fatal(err)
	}
	azlsPrint(prs, cfg.CompanyName)
}

func init() {
	listCmd.AddCommand(azlsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

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
