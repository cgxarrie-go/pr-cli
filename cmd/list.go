package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/hashicorp/go-multierror"

	"github.com/cgxarrie-go/prq/cmd/azure"
	"github.com/cgxarrie-go/prq/cmd/github"
	"github.com/cgxarrie-go/prq/domain/models"
	"github.com/cgxarrie-go/prq/utils"
)

const (
	prIDColLength      int = 8
	prTitleColLength   int = 70
	prCreatedColLength int = 25
	prStatusColLength  int = 10
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "list PRs",
	Long:    `List Pull Requests from the specified provider according to config`,
	RunE: func(cmd *cobra.Command, args []string) error {

		gitOrigins, err := utils.GitOrigins(".")
		if err != nil {
			return err
		}

		azureOrigins := utils.Origins{}
		githubOrigins := utils.Origins{}

		for _, origin := range gitOrigins {
			if origin.IsAzure() {
				azureOrigins = azureOrigins.Append(origin)
			}
			if origin.IsGithub() {
				githubOrigins = githubOrigins.Append(origin)
			}
		}

		prs :=[]models.PullRequest{}

		if len(azureOrigins) > 0 {
			azPrs, azErr := azure.RunListCmd(cmd, azureOrigins)
			if azErr != nil {
				multierror.Append(err, azErr)
			}
			prs = append(prs, azPrs...)
		}

		if len(githubOrigins) > 0 {
			ghPrs, ghErr := github.RunListCmd(cmd, githubOrigins)
			if ghErr != nil {
				multierror.Append(err, ghErr)
			}
			prs = append(prs, ghPrs...)
		}

		if len(prs) > 0 {
			printList(prs)
		}

		return err

	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

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

		created := fmt.Sprintf("%s (%v-%d-%d)",
			strings.Split(pr.CreatedBy, " ")[0],
			pr.Created.Year(), pr.Created.Month(), pr.Created.Day())
		format := getColumnFormat()
		title := pr.ShortenedTitle(70)
		prInfo := fmt.Sprintf(format, pr.ID, title,
			created, pr.Status, pr.Link)
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
