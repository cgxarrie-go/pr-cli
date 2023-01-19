package main

import (
	"fmt"
	"log"

	"github.com/cgxarrie/pr-go/domain/azure"
)

func main() {

	svc := azure.NewAzureService("companyNamw", "azurePAT")

	req := azure.GetPRsRequest{
		ProjectID:    "projectID",
		RepositoryID: "repoID",
		Status:       1,
	}

	prs, err := svc.GetPRs(req)
	if err != nil {
		log.Fatal(err)
	}

	for _, pr := range prs {
		fmt.Printf("%+v \n", pr)
	}
}
