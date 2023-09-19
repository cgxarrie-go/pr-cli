package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/cmd/azure"
	"github.com/cgxarrie-go/prq/cmd/github"
	"github.com/cgxarrie-go/prq/internal/models"
	"github.com/cgxarrie-go/prq/internal/utils"
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

		origin, err := utils.NewOrigin(".")
		if err != nil {
			return err
		}
		origins := utils.Origins{
			origin,
		}

		if origin.IsAzure() {
			prs, err := azure.RunListCmd(origins)
			if err != nil {
				return err
			}
			printList(prs)
			return nil
		}
		if origin.IsGithub() {
			prs, err := github.RunListCmd(origins)
			if err != nil {
				return err
			}
			printList(prs)
			return nil
		}

		return nil
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

	sort.SliceStable(prs, func(i,j int) bool {
		if prs[i].Project.Name != prs[j].Project.Name {
			return prs[i].Project.Name < prs[j].Project.Name
		}

		if prs[i].Repository.Name != prs[j].Repository.Name {
			return prs[i].Repository.Name < prs[j].Repository.Name
		}

		return prs[i].ID < prs[j].ID
	})
	

	for i, pr := range prs {
		if pr.Project.ID != lastProject {
			fmt.Println(prinTableTitle(pr.Project.Name))
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

func prinTableTitle(projectName string) string {

	format := "%s: %s\n    %s\n" + getColumnFormat()
	head := fmt.Sprintf(format, "Project", projectName, "Repository", "ID",
		"Title", "Created By", "Status", "Link")
	line := strings.Repeat("-", len(head)+5)
	doubleLine := strings.Repeat("=", len(head)+5)

	return fmt.Sprintf("%s\n%s\n%s", doubleLine, head, line)
}

func getColumnFormat() string {
	return "        %" + fmt.Sprintf("%d", prIDColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prTitleColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prCreatedColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prStatusColLength) + "s " +
		"| %s"
}
