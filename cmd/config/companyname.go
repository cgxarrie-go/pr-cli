package config

import (
	"fmt"

	appCfg "github.com/cgxarrie-go/prcli/config"
	"github.com/cgxarrie-go/prcli/domain/errors"
	"github.com/spf13/cobra"
)

// companyNameCmd represents the companyName command
var companyNameCmd = &cobra.Command{
	Use:   "company-name",
	Short: "set azure company name",
	Long:  `Set the Azure Conpmany-Name in the configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		err := runCompanyNameCmd(args)
		if err != nil {
			errors.Print(err)
			return
		}
	},
}

func init() {
	ConfigCmd.AddCommand(companyNameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// companyNameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// companyNameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runCompanyNameCmd(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid number of arguments")
	}
	cfg := appCfg.GetInstance()
	cfg.Load()

	cfg.Azure.CompanyName = args[0]
	err := cfg.Save()
	if err != nil {
		return err
	}
	return nil
}
