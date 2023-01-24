package main

import (
	"fmt"
	"log"

	"github.com/cgxarrie/pr-go/domain/azure"
)

func main() {

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
		Status: 1,
	}

	prs, err := svc.GetPRs(req)
	if err != nil {
		log.Fatal(err)
	}

	for _, pr := range prs {
		fmt.Printf("%+v \n", pr)
	}
}
