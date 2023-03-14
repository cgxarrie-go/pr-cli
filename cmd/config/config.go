package config

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	appcfg "github.com/cgxarrie-go/prq/config"
)

// ConfigCmd represents the Config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "display config",
	Long:  `display config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runConfigCmd(cmd, args)
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

}

func runConfigCmd(cmd *cobra.Command, args []string) error {

	cfg := appcfg.GetInstance()
	cfg.Load()
	b, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
