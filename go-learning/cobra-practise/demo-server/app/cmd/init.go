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

var _ phases.InitData = &initData{}

type initData struct {
	dryRun          bool
	aliceConfigDir  string
	aliceConfigPath string
	externalInitCfg string
	version         string
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
			c, err := initRunner.InitData(args)
			if err != nil {
				return err
			}

			data := c.(*initData)
			fmt.Printf("[init] Using Alice version: %s\n", data.Version())

			return initRunner.Run(args)
		},
		Args: cobra.NoArgs,
	}

	AddInitConfigFlags(cmd.Flags(), initOptions)

	initRunner.AppendPhase(phases.NewCertsPhase())
	initRunner.AppendPhase(phases.NewStartPhase())

	initRunner.SetDataInitializer(func(cmd *cobra.Command, args []string) (workflow.RunData, error) {
		data := newInitData(cmd, args, initOptions)
		return data, nil
	})

	initRunner.BindToCommand(cmd)

	return cmd
}

func AddInitConfigFlags(flagSet *flag.FlagSet, initOptions *initOptions) {
	flagSet.StringVar(&initOptions.externalInitCfg, "configFile", initOptions.externalInitCfg, "configFile Path")
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

func newInitData(cmd *cobra.Command, args []string, options *initOptions) *initData {
	return &initData{
		dryRun:          true,
		aliceConfigDir:  options.aliceConfigDir,
		aliceConfigPath: options.aliceConfigPath,
		externalInitCfg: options.externalInitCfg,
		version:         "v1.0",
	}

}

func (d *initData) DryRun() bool {
	return d.dryRun
}

func (d *initData) ConfigDir() string {
	return d.aliceConfigDir
}

func (d *initData) ConfigPath() string {
	return d.aliceConfigPath
}

func (d *initData) ExternalInitCfg() string {
	return d.externalInitCfg
}

func (d *initData) Version() string {
	return d.version
}
