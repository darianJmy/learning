package cmd

import (
	"github.com/spf13/cobra"
)

func NewDefaultAliceCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:  "alice",
		Long: "alice: easy practise cobra frame",
		//设置为 true 时可以在命令执行过程中遇到任何错误时，不显示错误
		SilenceErrors: true,
		//设置为, true 时可以在命令执行遇到输入错误时，不显示使用方法
		SilenceUsage: true,
	}

	cmd.ResetFlags()

	cmd.AddCommand(newCmdInit(nil))
	cmd.AddCommand(newCmdVersion())

	return cmd
}
