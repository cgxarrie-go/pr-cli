/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/cgxarrie/pr-go/cmd"
	"github.com/cgxarrie/pr-go/config"
)

func main() {
	cfg := config.GetInstance()
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
