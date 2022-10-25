package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func newCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of alice",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunVersion()
		},
	}
	return cmd
}

func RunVersion() error {
	klog.V(1).Infoln("[version] retrieving version info")
	fmt.Println("v1.0")
	return nil
}
