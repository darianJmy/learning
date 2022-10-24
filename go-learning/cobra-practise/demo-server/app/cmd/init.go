package cmd

import (
	"fmt"
	phases "github.com/darianJmy/learning/go-learning/cobra-practise/demo-server/app/cmd/phases/init"
	"github.com/darianJmy/learning/go-learning/cobra-practise/demo-server/app/cmd/phases/workflow"
	"path/filepath"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

type initOptions struct {
	dryRun          bool
	aliceConfigDir  string
	aliceConfigPath string
	externalInitCfg string
}

//var _ phases.InitData = &initData{}

type initData struct {
	dryRun          bool
	aliceConfigDir  string
	aliceConfigPath string
	externalInitCfg string
}

func newCmdInit(initOptions *initOptions) *cobra.Command {
	if initOptions == nil {
		initOptions = newInitOptions()
	}
	initRunner := workflow.NewRunner()

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Run this command in order to set up the Alice control plane",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf(initOptions.aliceConfigPath)
			//c, err := initRunner.InitData(args)
			//if err != nil {
			//	return err
			//}
			//
			//data := c.(*initData)
			//fmt.Printf("[init] Using Alice version: %s\n", "1.0")
			//
			//if err = initRunner.Run(args); err != nil {
			//	return err
			//}
			//return showJoinCommand(data)
			return nil
		},
		Args: cobra.NoArgs,
	}

	AddInitConfigFlags(cmd.Flags(), initOptions)

	initRunner.AppendPhase(phases.NewCertsPhase())
	return cmd
}

func AddInitConfigFlags(flagSet *flag.FlagSet, initOptions *initOptions) {
	flagSet.StringVar(&initOptions.externalInitCfg, "configFile", initOptions.externalInitCfg, "configFile Path")
}

func RunInit(initOptions *initOptions) error {
	fmt.Printf("Hello world, %v", initOptions.externalInitCfg)
	return nil
}

func newInitOptions() *initOptions {
	return &initOptions{
		aliceConfigDir:  "/etc/alice",
		aliceConfigPath: GetAdminAliceConfigPath(),
	}
}

func GetAdminAliceConfigPath() string {
	return filepath.Join("/etc/alice", "admin.conf")
}
