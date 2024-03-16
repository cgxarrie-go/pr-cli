package cmd

import (
	"fmt"

	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/remoteclient"
	"github.com/cgxarrie-go/prq/internal/services"
	"github.com/cgxarrie-go/prq/internal/utils"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

// ListCmd represents the list command
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Create Pull Request",
	Long:    "Create a Pull Request.",
	RunE: func(cmd *cobra.Command, args []string) error {
		dest, _ := cmd.Flags().GetString("target")
		ttl, _ := cmd.Flags().GetString("title")
		dft, _ := cmd.Flags().GetString("draft")
		draft := !utils.IsFalse(dft)

		r, err := utils.CurrentFolderRemote()
		if err != nil {
			return fmt.Errorf("getting remote: %w", err)
		}

		cl, err := remoteclient.NewRemoteClient(r)
		if err != nil {
			return fmt.Errorf("creating remote client: %w", err)
		}

		svc := services.NewCreatePRService(cl)
		svcReq := ports.CreatePRSvcRequest{
			Destination: dest,
			Title:       ttl,
			IsDraft:     draft,
		}

		pr, err := svc.Run(svcReq)
		if err != nil {
			return fmt.Errorf("creating PR: %w", err)
		}

		lnk := termenv.Hyperlink(pr.Link, "open PR")

		fmt.Printf("PR created with ID: %s (%s)\n", pr.ID, lnk)
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

	createCmd.Flags().StringP("target", "b", "", "target branch. If blank, master is used")
	createCmd.Flags().StringP("title", "t", "", "title. If blank, standard title is used")
	createCmd.Flags().StringP("draft", "d", "", "draft. default is true")

}
