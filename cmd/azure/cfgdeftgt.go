package azure

import (
	"fmt"

	"github.com/cgxarrie-go/prq/internal/config"
	"github.com/spf13/cobra"
)

// CfgDefTgtBranchCmd config azure default target branch command
var CfgDefTgtBranchCmd = &cobra.Command{
	Use:   "aztgt",
	Short: "set azure default target branch",
	Long:  `Set the Azure default target branch in the configuration file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runCfgDefTgtBranchCmd(args)
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

func runCfgDefTgtBranchCmd(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid number of arguments")
	}
	cfg := config.GetInstance()
	cfg.Load()

	cfg.Azure.DefaultTargetBranch = args[0]
	err := cfg.Save()
	if err != nil {
		return err
	}
	return nil
}
