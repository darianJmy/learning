package phases

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/darianJmy/learning/go-learning/cobra-practise/demo-server/app/cmd/phases/workflow"
)

func NewStartPhase() workflow.Phase {
	return workflow.Phase{
		Name:  "start",
		Short: "Start generation",
		Long:  "This command is not meant to be run on its own. See list of available subcommands.",
		Run:   runStart,
	}
}

func runStart(c workflow.RunData) error {
	data, ok := c.(InitData)
	if !ok {
		return errors.New("certs phase invoked with an invalid data struct")
	}

	fmt.Printf("[start] Using start service alice\n")
	fmt.Println(data.ExternalInitCfg())
	return nil
}
