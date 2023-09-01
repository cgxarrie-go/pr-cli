package create

import (
	"github.com/spf13/cobra"
)

// ListCmd represents the list command
var CreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"t"},
	Short:   "Create Pull Request",
	Long:    `Create a Pull Requests in the the specified provider according to config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		prj, _ := cmd.Flags().GetString("project")
		repo, _ := cmd.Flags().GetString("repository")
		dest, _ := cmd.Flags().GetString("destination")
		src, _ := cmd.Flags().GetString("source")
		ttl, _ := cmd.Flags().GetString("title")

		err := runCreateAzureCmd(cmd, prj, repo, src, dest, ttl)
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

	CreateCmd.Flags().StringP("project", "p", "", "project")
	CreateCmd.Flags().StringP("repository", "r", "", "repository where the PR is to be created")
	CreateCmd.Flags().StringP("source", "s", "", "source branch")
	CreateCmd.Flags().StringP("destination", "d", "", "target branch")
	CreateCmd.Flags().StringP("title", "t", "", "title")

}
