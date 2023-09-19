package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/cmd/azure"
	"github.com/cgxarrie-go/prq/cmd/github"
	"github.com/cgxarrie-go/prq/internal/config"
)

// configCmd represents the Config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "display config",
	Long:  `display config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runConfigCmd(cmd, args)
	},
}

func init() {
	configCmd.AddCommand(azure.CfgPATCmd)
	configCmd.AddCommand(github.CfgPATCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func runConfigCmd(cmd *cobra.Command, args []string) error {

	cfg := config.GetInstance()
	err := cfg.Load()
	if err != nil {
		return fmt.Errorf("reading config: %w", err)
	}
	b, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("marshalling config: %w", err)
	}
	fmt.Println(string(b))
	return nil
}
