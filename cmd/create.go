package cmd

import (
	"fmt"

	"github.com/cgxarrie-go/prq/cmd/azure"
	"github.com/cgxarrie-go/prq/cmd/github"
	"github.com/cgxarrie-go/prq/utils"
	"github.com/spf13/cobra"
)

// ListCmd represents the list command
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Create Pull Request",
	Long: "Create a Pull Request." +
		"\n\tSource branch is the current repository active brnach" +
		"\n\tDestination branch can be specified by flag -d. If ommitted, destination will be master" +
		"\n\tTitle can be specified by flag -t. If ommitted, title will be standard title",
	RunE: func(cmd *cobra.Command, args []string) error {
		dest, _ := cmd.Flags().GetString("destination")
		ttl, _ := cmd.Flags().GetString("title")

		o, err := utils.CurrentOrigin()
		if err != nil {
			return fmt.Errorf("getting origin: %w", err)
		}
		if o.IsAzure() {
			err := azure.RunCreatCmd(cmd, dest, ttl)	
			return err
		}

		if o.IsGithub() {
			err := github.RunCreatCmd(cmd, dest, ttl)
			return err
		}

		return fmt.Errorf("unknown origin type: %s", o)
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

	createCmd.Flags().StringP("destination", "d", "", "target branch. If blank, master is used")
	createCmd.Flags().StringP("title", "t", "", "title. If blank, standard title is used")

}
