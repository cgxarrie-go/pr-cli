package cmd

import (
	"github.com/cgxarrie-go/prq/internal/config"
	"github.com/spf13/cobra"
)

var configGithubCmd = &cobra.Command{
	Use:   "az",
	Short: "config Github",
	Long:  `config Github`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pat, _ := cmd.Flags().GetString("pat")
		branch, _ := cmd.Flags().GetString("branch")
		err := runConfigGithubCmd(pat, branch)
		return err
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GithubPATCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	configGithubCmd.Flags().StringP("pat", "p", "", "set the PAT for Github DevOps")
	configGithubCmd.Flags().StringP("branch", "b", "", "set the default target branch for creation of PRs")
}

func runConfigGithubCmd(pat string, branch string) error {

	if pat == "" && branch == "" {
		return nil
	}

	cfg := config.GetInstance()
	cfg.Load()

	if pat != "" {
		cfg.Github.PAT = pat
	}

	if branch != "" {
		cfg.Github.DefaultTargetBranch = branch
	}

	err := cfg.Save()
	if err != nil {
		return err
	}
	return nil
}
