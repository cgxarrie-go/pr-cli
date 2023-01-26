package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/cgxarrie/pr-go/config"
	"github.com/cgxarrie/pr-go/domain/models"
	"github.com/cgxarrie/pr-go/services/azure"
	"github.com/spf13/cobra"
)

// azlsCmd represents the list command
var azlsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List PRs in azure",
	Long:  `List PRs in azure`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 || args[0] == "" {
			fmt.Println("Error : expected one argument")
			return
		}

		status := 1
		switch args[0] {
		case "ac":
			status = 1
		case "ab":
			status = 2
		case "co":
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

		req := azure.GetPRsRequest{
			ProjectRepos: projectRepos,
			Status:       status,
		}

		prs, err := svc.GetPRs(req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Number of PRs : %d \n", len(prs))
		for i, pr := range prs {
			if i == 0 {
				fmt.Println(printableTitle())
			}
			url := fmt.Sprintf("https://dev.azure.com/%s/%s/_git/%s/"+
				"pullrequest/%s", cfg.CompanyName, pr.ProjectName,
				pr.RepositoryName, pr.ID)
			fmt.Println(printable(pr, url))
		}
	},
}

func init() {
	azCmd.AddCommand(azlsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func printable(pr models.PullRequest, url string) string {

	return fmt.Sprintf("%8s | %-20s | %-15s | %-50s | %-20s | %-10s | %s", pr.ID, pr.ProjectName, pr.RepositoryName,
		pr.Title, pr.CreatedBy, pr.Status, url)
}

func printableTitle() string {

	head := fmt.Sprintf("%8s | %-20s | %-15s | %-50s | %-20s | %-10s | %s", "ID", "Project", "Repository",
		"Title", "Created By", "Status", "URL")
	line := strings.Repeat("-", len(head)+5)

	return fmt.Sprintf("%s\n%s", head, line)
}
