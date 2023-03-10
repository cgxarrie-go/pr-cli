package cmd

import (
	"os"

	"github.com/cgxarrie-go/prcli/cmd/config"
	"github.com/cgxarrie-go/prcli/cmd/list"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "prq",
	Short: "Interaction with pull requests from command line",
	Long: `Interaction with pull requests from command line
	
Config commands : 
config : display config
config az-cname	: set company name in azure config
config az-pat	: set PAT in azure config
config az-project -a <name>	: Adds a project with name <name> in azure config
config az-project -d <name>	: Removes a project with name <name> in azure config
config az-repo -a <name> -p <project-name>	: Adds a repo with name <name> to the project with name <project-name> in azure config
config az-repo -d <name> -p <project-name>	: Removes a repo with name <name> from the project with name <project-name> in azure config

List PR commands : 
list az : Lists all PR in status Active for azure projects and repos
list az active: Lists all PR in status Active for azure projects and repos
list az abandoned: Lists all PR in status Abandoned for azure projects and repos
list az cancelled: Lists all PR in status Cancelled for azure projects and repos
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(config.ConfigCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pr-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
