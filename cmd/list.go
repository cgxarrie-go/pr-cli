package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/muesli/termenv"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/internal/config"
	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/remote"
	"github.com/cgxarrie-go/prq/internal/remoteclient"
	"github.com/cgxarrie-go/prq/internal/remotetype"
	"github.com/cgxarrie-go/prq/internal/services"
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
	Long: `List Pull Requests from the specified provider according to config
	
	Flags
	-o, -- option	: option to list PRs from current directory tree (d), config (c) or current directory (default)
	-f, -- filter	: filter PRs by any field`,

	RunE: func(cmd *cobra.Command, args []string) error {
		filter, _ := cmd.Flags().GetString("filter")
		opt, _ := cmd.Flags().GetString("option")

		var remotes remote.Remotes
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
			remotes = make(remote.Remotes, len(cfg.Remotes))
			for _, c := range cfg.Remotes {
				r, _ := remote.NewRemote(c)
				remotes[r] = struct{}{}
			}

		default:
			currentRemote, err := utils.CurrentFolderRemote()
			if err != nil {
				return errors.Wrapf(err, "getting remote from current directory")
			}
			remotes = remote.Remotes{
				currentRemote: struct{}{},
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

func runListCmd(remotes remote.Remotes, filter string) error {

	config.GetInstance().Load()

	unknownRemotes := remote.Remotes{}

	for r := range remotes {
		switch r.Type() {
		case remotetype.NotSet:
			unknownRemotes.Append(r)
		default:
			cl, _ := remoteclient.NewRemoteClient(r)
			svc := services.NewGetPRsService(cl)
			resp := svc.Run()
			printRemoteHeader(resp.Remote, resp.Count)

			if resp.Error != nil {
				fmt.Println(resp.Error)
				continue
			}

			printList(resp, filter)
		}
	}
	return nil

}

func printList(req ports.GetPRsSvcResponse, filter string) {
	filter = strings.ToLower(filter)

	tableTitle := getTableTitle()
	fmt.Println(tableTitle)
	fmt.Println(strings.Repeat("-", len(tableTitle)+5))

	sort.SliceStable(req.PullRequests, func(i, j int) bool {
		return req.PullRequests[i].Created.Before(
			req.PullRequests[j].Created)
	})

	for _, pr := range req.PullRequests {

		status := "Active"
		if pr.IsDraft {
			status = "Draft"
		}
		created := fmt.Sprintf("%s (%v-%d-%d)",
			strings.Split(pr.CreatedBy, " ")[0],
			pr.Created.Year(), pr.Created.Month(), pr.Created.Day())
		format := getColumnFormat()
		title := shortenedText(pr.Title, 70)
		lnk := termenv.Hyperlink(pr.Link, "open PR")
		prInfo := fmt.Sprintf(format, pr.ID, status, title,
			created, lnk)

		if filter != "" && !strings.Contains(strings.ToLower(prInfo), filter) {
			continue
		}

		fmt.Println(prInfo)
	}
}

func getTableTitle() string {
	format := getColumnFormat()
	return fmt.Sprintf(format, "ID", "Status", "Title", "Created By", "Link")
}

func printRemoteHeader(remote string, count int) {

	head := getTableTitle()

	line := strings.Repeat("-", len(head)+5)
	doubleLine := strings.Repeat("=", len(head)+5)

	fmt.Println()
	fmt.Println(doubleLine)
	fmt.Printf("Remote: %s\n", remote)
	fmt.Printf("Number of PRs: %d\n", count)
	fmt.Println(line)
}

func shortenedText(text string, maxLength int) string {

	pritntable := text

	if len(pritntable) <= maxLength {
		return pritntable
	}

	shortenLenght := maxLength - 3

	title := fmt.Sprintf("%s...", pritntable[0:shortenLenght])
	return title
}

func getColumnFormat() string {
	return "%" + fmt.Sprintf("%d", prIDColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prStatusColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prTitleColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", prCreatedColLength) + "s " +
		"| %s"
}
