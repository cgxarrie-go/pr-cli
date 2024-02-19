package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/cmd/azure"
	"github.com/cgxarrie-go/prq/cmd/github"
	"github.com/cgxarrie-go/prq/internal/config"
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
		opt, _ := cmd.Flags().GetString("option")

		var remotes utils.Remotes
		var err error

		switch opt {
		case "d":
			remotes, err = utils.CurrentFolderTreeRemotes()
			if err != nil {
				return errors.Wrapf(err, "getting remotes from current directory tree")
			}
		case "c":
			cfg := config.GetInstance()
			err := cfg.Load()
			if err != nil {
				return errors.Wrapf(err, "getting remotes from config")
			}
			remotes = make(utils.Remotes, len(cfg.Remotes))
			for i, c := range cfg.Remotes {
				remotes[i] = utils.Remote(c)
			}

		default:
			currentRemote, err := utils.CurrentFolderRemote()
			if err != nil {
				return errors.Wrapf(err, "getting remote from current directory")
			}
			remotes = utils.Remotes{
				currentRemote,
			}
		}

		return runListCmd(remotes)

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

	listCmd.Flags().StringP("option", "o", "", "option")
}

func runListCmd(remotes utils.Remotes) error {

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
		fmt.Println(msg)
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

}

func printList(prs []models.PullRequest) {
	fmt.Printf("Number of PRs : %d \n", len(prs))
	lastOrigin := ""

	sort.SliceStable(prs, func(i, j int) bool {
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

		status := pr.Status
		if pr.IsDraft {
			status = "Draft"
		}
		created := fmt.Sprintf("%s (%v-%d-%d)",
			strings.Split(pr.CreatedBy, " ")[0],
			pr.Created.Year(), pr.Created.Month(), pr.Created.Day())
		format := getColumnFormat()
		title := pr.ShortenedTitle(70)
		prInfo := fmt.Sprintf(format, pr.ID, status, title,
			created, pr.Link)
		fmt.Println(prInfo)
	}
}

func prinTableTitle(remote string) string {

	format := "%s\n" + getColumnFormat()
	head := fmt.Sprintf(format, remote, "ID", "Status",
		"Title", "Created By", "Link")
	line := strings.Repeat("-", len(head)+5)
	doubleLine := strings.Repeat("=", len(head)+5)

	return fmt.Sprintf("%s\n%s\n%s", doubleLine, head, line)
}

func getColumnFormat() string {
	return "%" + fmt.Sprintf("%d", prIDColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prStatusColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prTitleColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prCreatedColLength) + "s " +
		"| %s"
}
