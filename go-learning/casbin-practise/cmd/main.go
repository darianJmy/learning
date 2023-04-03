package main

import (
	"log"
	"os"

	"casbin-practise/cmd/app"
)

func main() {
	cmd := app.NewDefaultCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
