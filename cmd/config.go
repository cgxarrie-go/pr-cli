/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prcli/config"
)

// listCmd represents the az command
var configComd = &cobra.Command{
	Use:   "config",
	Short: "display config",
	Long:  `display config`,
	Run: func(cmd *cobra.Command, args []string) {
		runConfigCmd(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(configComd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func runConfigCmd(cmd *cobra.Command, args []string) {

	cfg := config.GetInstance().Azure
	b, err := json.Marshal(cfg)
	if err != nil {
		fmt.Printf("ERROR : %s\n", err.Error())
		return
	}
	fmt.Println(string(b))
}
