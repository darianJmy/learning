package cmd

import (
	options2 "github.com/darianJmy/learning/go-learning/casbin-practise/app/cmd/options"
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/core"
)

var CoreV1 core.CoreV1Interface

// Setup 完成核心应用接口的设置
func Setup(o *options2.Options) {
	CoreV1 = core.New(o.Factory)
}
