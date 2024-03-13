package cmd

import (
	"fmt"

	"github.com/cgxarrie-go/prq/cmd/azure"
	"github.com/cgxarrie-go/prq/cmd/github"
	"github.com/cgxarrie-go/prq/internal/remote"
	"github.com/cgxarrie-go/prq/internal/utils"
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

		o, err := remote.CurrentFolderRemote()
		if err != nil {
			return fmt.Errorf("getting origin: %w", err)
		}

		switch o.Type() {
		case remote.RemoteTypeGitHub:
			err := github.RunCreatCmd(cmd, dest, ttl, draft)
			return err
		case remote.RemoteTypeAzure:
			err := azure.RunCreatCmd(cmd, dest, ttl, draft)
			return err
		default:
			return fmt.Errorf("unknown origin type: %s", o)

		}
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
