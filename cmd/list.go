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
		filter, _ := cmd.Flags().GetString("filter")
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

		return runListCmd(remotes, filter)

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
	listCmd.Flags().StringP("filter", "f", "", "")
}

func runListCmd(remotes utils.Remotes, filter string) error {

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

	printList(prs, filter)
	return nil

}

func printList(prs []models.PullRequest, filter string) {
	filter = strings.ToLower(filter)

	sort.SliceStable(prs, func(i, j int) bool {
		if prs[i].Origin != prs[j].Origin {
			return prs[i].Organization < prs[j].Origin
		}

		return prs[i].Created.Before(prs[j].Created)
	})

	lastOrigin := ""
	visiblePRs := []string{}

	for _, pr := range prs {
		if pr.Origin != lastOrigin {
			if len(visiblePRs) > 0 {
				printRemotePRs(lastOrigin, visiblePRs)
			}
			lastOrigin = pr.Origin
			visiblePRs = []string{}
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

		if filter != "" && !strings.Contains(strings.ToLower(prInfo), filter) {
			continue
		}

		visiblePRs = append(visiblePRs, prInfo)

	}
	printRemotePRs(lastOrigin, visiblePRs)
}

func printRemoteHeader(remote string, count int) {

	format := getColumnFormat()
	head := fmt.Sprintf(format, "ID", "Status", "Title", "Created By", "Link")

	line := strings.Repeat("-", len(head)+5)
	doubleLine := strings.Repeat("=", len(head)+5)

	fmt.Println()
	fmt.Println(doubleLine)
	fmt.Printf("Remote: %s\n", remote)
	fmt.Printf("Number of PRs: %d\n", count)
	fmt.Println(line)
}

func getColumnFormat() string {
	return "%" + fmt.Sprintf("%d", prIDColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prStatusColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prTitleColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prCreatedColLength) + "s " +
		"| %s"
}

func printRemotePRs(remote string, prs []string) {

	printRemoteHeader(remote, len(prs))
	for _, pr := range prs {
		fmt.Println(pr)
	}
}
