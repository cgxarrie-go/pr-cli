/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cgxarrie/pr-go/domain/azure"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List active PRs",
	Long: `List all Pull Requests in all the projects and repositores
	stated in the config file azure section`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 || args[0] == "" {
			fmt.Println("Error : expected one argument")
			return
		}

		status, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error Invalid argument: %s", args[0])
			return
		}

		svc := azure.NewAzureService("companyName", "azurePAT")
		req := azure.GetPRsRequest{
			ProjectRepos: map[string][]string{
				"proj.A": {
					"Repo.A.1",
					"Repo.A.2",
					"Repo.A.3",
				},
				"proj.B": {
					"Repo.B.1",
					"Repo.B.2",
					"Repo.B.3",
				},
			},
			Status: status,
		}

		prs, err := svc.GetPRs(req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Number of PRs : %d \n", len(prs))
		for _, pr := range prs {
			fmt.Printf("%+v \n", pr)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
