package config

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/cache/providers"
	appCfg "github.com/cgxarrie-go/prq/config"
)

// azureCmd represents the azure configuration command
var defaultProviderCommandCmd = &cobra.Command{
	Use:     "default-provider",
	Aliases: []string{"def-prov", "dp"},
	Short:   "sets the default provider for future commands",
	Long:    `sets the default provider for future commands`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runDefaultProviderCmd(args)
		return err
	},
}

func init() {
	ConfigCmd.AddCommand(defaultProviderCommandCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azurePATCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azurePATCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runDefaultProviderCmd(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid number of arguments")
	}
	cfg := appCfg.GetInstance()
	cfg.Load()

	provider, err := providers.FromName(args[0])
	if err != nil {
		return err
	}

	cfg.DefaultProvider = provider
	err = cfg.Save()
	if err != nil {
		return err
	}
	return nil
}
