package init

import "github.com/darianJmy/learning/go-learning/cobra-practise/demo-server/app/cmd/phases/workflow"

func NewCertsPhase() workflow.Phase {
	return workflow.Phase{
		Name:  "certs",
		Short: "Certificate generation",
		Long:  "This command is not meant to be run on its own. See list of available subcommands.",
		//Phases: newCertSubPhases(),
		Run: runCerts,
	}
}

//func newCertSubPhases() []workflow.Phase {
//	subPhases := []workflow.Phase{}
//
//	allPhase := workflow.Phase{
//		Name:           "all",
//		Short:          "Generate all certificates",
//		InheritFlags:   getCertPhaseFlags("all"),
//		RunAllSiblings: true,
//	}
//
//	subPhases = append(subPhases, allPhase)
//
//	return subPhases
//}
//
//func getCertPhaseFlags(name string) []string {
//	flags := []string{
//		"2131", "213asd", "ssda",
//	}
//	if name == "all" || name == "apiserver" {
//		flags = append(flags, "123")
//	}
//
//	return flags
//}

func runCerts() error {

	//data, ok := c.(InitData)
	//if !ok {
	//	return errors.New()
	//}

}
