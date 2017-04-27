package main

import (
	"log"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd"
)

func main() {
	alauda, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	rootCmd := cmd.NewRootCmd(alauda)

	err = rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
