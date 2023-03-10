package list

import (
	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/services/azure/status"
)

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "list PRs",
	Long:    `List Pull Requests from the specified provider according to config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		st, _ := cmd.Flags().GetString("status")
		if st == "" {
			st = status.Active.Name()
		}
		err := runListAzureCmd(cmd, st)
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
	ListCmd.Flags().StringP("status", "s", status.Active.Name(), "status of PRs to list")

}
