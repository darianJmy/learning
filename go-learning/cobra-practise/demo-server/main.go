package main

import (
	"os"

	"github.com/darianJmy/learning/go-learning/cobra-practise/demo-server/app/cmd"
)

func main() {
	cmd := cmd.NewDefaultAliceCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
