package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
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

		currentRemote, err := utils.CurrentFolderRemote()
		if err != nil {
			return errors.Wrapf(err, "getting remote from current directory")
		}
		remotes := utils.Remotes{
			currentRemote,
		}

		azRemotes := utils.Remotes{}
		ghRemotes := utils.Remotes{}
		unknownRemotes := utils.Remotes{}

		prs := []models.PullRequest{}	
		for _, remote := range remotes {
			switch true {
			case remote.IsAzure():
				azRemotes.Append(remote)
			case remote.IsGithub():
				ghRemotes.Append(remote)
			default:
				unknownRemotes.Append(remote)
			}
		}

		if len(unknownRemotes) > 0 {
			msg := ""
			for _, ur := range unknownRemotes {
				msg = fmt.Sprintf("%s%s/n", msg, ur)
			}
			msg = fmt.Sprintf("unknown remote types/n%s", msg)
		}

		if len(azRemotes) > 0 {
			azPrs, err := azure.RunListCmd(azRemotes)
			if err != nil {
				return errors.Wrapf(err, "getting PRs from azure repositories")
			}
			prs = append(prs, azPrs...)
		}

		if len(ghRemotes) > 0 {
			ghPrs, err := github.RunListCmd(ghRemotes)
			if err != nil {
				return errors.Wrapf(err, "getting PRs from azure github")
			}
			prs = append(prs, ghPrs...)
		}

		printList(prs)			
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
	lastOrigin := ""

	sort.SliceStable(prs, func(i,j int) bool {
		if prs[i].Origin != prs[j].Origin {
			return prs[i].Organization < prs[j].Origin
		}

		return prs[i].ID < prs[j].ID
	})
	

	for _, pr := range prs {
		if pr.Origin != lastOrigin {
			fmt.Println(prinTableTitle(pr.Origin))
			lastOrigin = pr.Origin
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

func prinTableTitle(remote string) string {

	format := "%s\n" + getColumnFormat()
	head := fmt.Sprintf(format, remote, "ID",
		"Title", "Created By", "Status", "Link")
	line := strings.Repeat("-", len(head)+5)
	doubleLine := strings.Repeat("=", len(head)+5)

	return fmt.Sprintf("%s\n%s\n%s", doubleLine, head, line)
}

func getColumnFormat() string {
	return "%" + fmt.Sprintf("%d", prIDColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prTitleColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prCreatedColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prStatusColLength) + "s " +
		"| %s"
}
