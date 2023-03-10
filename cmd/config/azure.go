package config

import (
	"github.com/spf13/cobra"
)

// azureCmd represents the azure configuration command
var azureCmd = &cobra.Command{
	Use:     "azure",
	Aliases: []string{"az"},
	Short:   "set the Zzure configuration",
	Long:    `Set the Azure configuration`,
}

func init() {
	ConfigCmd.AddCommand(azureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azurePATCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azurePATCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
