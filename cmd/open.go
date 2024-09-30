package cmd

import (
	"fmt"

	"github.com/cgxarrie-go/prq/internal/config"
	"github.com/cgxarrie-go/prq/internal/remoteclient"
	"github.com/cgxarrie-go/prq/internal/utils"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:     "open",
	Aliases: []string{"o"},
	Short:   "Open Pull Request",
	Long:    `Open default browser to show the selected Pull Request`,
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

		if len(args) == 0 || args[0] == "" {
			return fmt.Errorf("PR ID is required")
		}

		id := args[0]
		err = cl.Open(id)
		if err != nil {
			return fmt.Errorf("opening browser for PR %s: %w", id, err)
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
