package cmd

import (
	"fmt"

	"github.com/cgxarrie-go/prq/internal/config"
	"github.com/cgxarrie-go/prq/internal/remoteclient"
	"github.com/cgxarrie-go/prq/internal/utils"
	"github.com/spf13/cobra"
)

var codeCmd = &cobra.Command{
	Use:     "code",
	Aliases: []string{"d"},
	Short:   "Open Repository code",
	Long:    `Open default browser to show the repository code`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config.GetInstance().Load()

		r, err := utils.CurrentFolderRemote()
		if err != nil {
			return fmt.Errorf("getting remote: %w", err)
		}

		cl, err := remoteclient.NewRemoteClient(r)
		if err != nil {
			return fmt.Errorf("creating remote client: %w", err)
		}

		err = cl.OpenCode()
		if err != nil {
			return fmt.Errorf("opening browser for repository %s: %w", r.Repository(), err)
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
