package github

import (
	"fmt"

	"github.com/cgxarrie-go/prq/internal/config"
	"github.com/spf13/cobra"
)

// CfgDefTgtBranchCmd config ghure default target branch command
var CfgDefTgtBranchCmd = &cobra.Command{
	Use:   "ghtgt",
	Short: "set Github default target branch",
	Long:  `Set the Github default target branch in the configuration file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runCfgDefTgtBranchCmd(args)
		return err
	},
}

func init() {

}

func runCfgDefTgtBranchCmd(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid number of arguments")
	}
	cfg := config.GetInstance()
	cfg.Load()

	cfg.Github.DefaultTargetBranch = args[0]
	err := cfg.Save()
	if err != nil {
		return err
	}
	return nil
}
