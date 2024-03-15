/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/cgxarrie-go/prq/cmd"
	"github.com/cgxarrie-go/prq/internal/github"
	"github.com/cgxarrie-go/prq/internal/services"
)

func main() {

	ghClient := github.NewClient("token")
	ghCreatePRSvc := services.NewCreatePRService("token", ghClient)

	cmd.Execute()
}
