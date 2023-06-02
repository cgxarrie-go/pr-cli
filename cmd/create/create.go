package create

import (
	"github.com/spf13/cobra"
)

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c", "cr"},
	Short:   "Create Pull Request",
	Long:    `Create a Pull Requests in the the specified provider according to config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repository")
		src, _ := cmd.Flags().GetString("source")
		tgt, _ := cmd.Flags().GetString("target")

		err := runCreateAzureCmd(cmd, repo, src, tgt)
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
	ListCmd.Flags().StringP("repository", "repo", "", "repository where the PR is to be created")
	ListCmd.Flags().StringP("source", "src", "", "source branch")
	ListCmd.Flags().StringP("target", "tgt", "", "target branch")

}
