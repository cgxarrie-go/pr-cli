package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "prq",
	Short: "Interaction with pull requests from command line",
	Long:  `Interaction with pull requests from command line`,
	Example: `Config commands : 
prq config : display config
prq config org : set company name in azure config
prq config pat : set PAT in azure config
prq config project -a <name> : Adds a project with name <name> in azure config
prq config project -d <name> : Removes a project with name <name> in azure config
prq config repo -p <project-name> -a <name> : Adds a repo with name <name> to the project with name <project-name> in azure config
prq config repo -p <project-name> -d <name> : Removes a repo with name <name> from the project with name <project-name> in azure config

List PR commands : 
prq list : Lists all PR in status Active for azure projects and repos
prq list --status active: Lists all PR in status Active for azure projects and repos
prq list --status abandoned: Lists all PR in status Abandoned for azure projects and repos
prq list --status cancelled: Lists all PR in status Cancelled for azure projects and repos`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pr-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
