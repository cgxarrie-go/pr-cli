package config

import (
	"fmt"

	"github.com/spf13/cobra"

	appCfg "github.com/cgxarrie-go/prq/config"
	"github.com/cgxarrie-go/prq/domain/errors"
)

// azureOrganizationCmd represents the organization command
var azureOrganizationCmd = &cobra.Command{
	Use:   "org",
	Short: "set azure company name",
	Long:  `Set the Azure Conpmany-Name in the configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		err := runOrganizationCmd(args)
		if err != nil {
			errors.Print(err)
			return
		}
	},
}

func init() {
	ConfigCmd.AddCommand(azureOrganizationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// organizationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// organizationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runOrganizationCmd(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid number of arguments")
	}
	cfg := appCfg.GetInstance()
	cfg.Load()

	cfg.Azure.Organization = args[0]
	err := cfg.Save()
	if err != nil {
		return err
	}
	return nil
}
