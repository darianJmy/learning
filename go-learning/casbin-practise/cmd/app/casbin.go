package app

import (
	"casbin-practise/pkg/global"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"casbin-practise/api"
	"casbin-practise/cmd/app/options"
)

type casbinOptions struct {
	configDir string
}

func newCmdRun(casbinOptions *casbinOptions) *cobra.Command {
	if casbinOptions == nil {
		casbinOptions = newCasbinOptions()
	}

	cmd := &cobra.Command{
		Use:   "run",
		Short: "This service operates with RBAC authorization management.",
		RunE: func(cmd *cobra.Command, args []string) error {
			opt := options.NewOptions(casbinOptions.configDir).
				NewContainer(restful.NewContainer())
			if err := opt.Complete(); err != nil {
				return err
			}
			return runCasbin(opt)
		},
	}

	addFlags(cmd.Flags(), casbinOptions)

	return cmd
}

func runCasbin(opt *options.Options) error {
	global.Setup(opt)

	api.RegisterRoutes(opt.Container)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", opt.ComponentConfig.Default.Listen),
		Handler: opt.Container,
	}

	log.Fatal(server.ListenAndServe())
	return nil
}

func addFlags(flagSet *flag.FlagSet, casbinOptions *casbinOptions) {
	flagSet.StringVar(&casbinOptions.configDir, "configFile", casbinOptions.configDir, "configFile Path")
}

func newCasbinOptions() *casbinOptions {
	return &casbinOptions{
		configDir: "/etc/project/config.yaml",
	}
}
