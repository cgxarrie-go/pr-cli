package config

import (
	"fmt"

	"github.com/spf13/cobra"

	appCfg "github.com/cgxarrie-go/prq/config"
)

// GithubPATCmd represents the azurePAT command
var GithubPATCmd = &cobra.Command{
	Use:   "ghpat",
	Short: "set Github PAT",
	Long:  `Set the Github PAT in the configuration file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runGithubPATCmd(args)
		return err
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azurePATCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azurePATCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runGithubPATCmd(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid number of arguments")
	}
	cfg := appCfg.GetInstance()
	cfg.Load()

	cfg.Github.PAT = args[0]
	err := cfg.Save()
	if err != nil {
		return err
	}
	return nil
}
