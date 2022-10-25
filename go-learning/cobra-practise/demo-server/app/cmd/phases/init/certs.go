package phases

import (
	"fmt"
	"github.com/darianJmy/learning/go-learning/cobra-practise/demo-server/app/cmd/phases/workflow"
	"github.com/pkg/errors"
)

func NewCertsPhase() workflow.Phase {
	return workflow.Phase{
		Name:   "certs",
		Short:  "Certificate generation",
		Long:   "This command is not meant to be run on its own. See list of available subcommands.",
		Phases: newCertSubPhases(),
		Run:    runCerts,
	}
}

func newCertSubPhases() []workflow.Phase {
	subPhases := []workflow.Phase{}

	allPhase := workflow.Phase{
		Name:           "all",
		Short:          "Generate all certificates",
		RunAllSiblings: true,
	}

	subPhases = append(subPhases, allPhase)

	for _, cert := range GetDefaultCertList() {
		var phase workflow.Phase
		phase = newCertSubPhase(cert, runCAPhase(cert))
		subPhases = append(subPhases, phase)

	}
	return subPhases
}

func newCertSubPhase(certSpec *Cert, run func(c workflow.RunData) error) workflow.Phase {
	phase := workflow.Phase{
		Name:  certSpec.Name,
		Short: fmt.Sprintf("Generate the %s", certSpec.LongName),
		Run:   run,
	}
	return phase
}

func runCAPhase(ca *Cert) func(c workflow.RunData) error {
	return func(c workflow.RunData) error {
		data, ok := c.(InitData)
		if !ok {
			return errors.New("certs phase invoked with an invalid data struct")
		}

		if !data.DryRun() {
			return errors.New("data")
		}

		fmt.Printf("phase name %s\n", ca.Name)

		return nil
	}
}

func runCerts(c workflow.RunData) error {
	data, ok := c.(InitData)
	if !ok {
		return errors.New("certs phase invoked with an invalid data struct")
	}

	fmt.Printf("[certs] Using certificateDir folder %q\n", data.ConfigDir())
	return nil
}
