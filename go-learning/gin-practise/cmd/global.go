package cmd

import (
	"github.com/darianJmy/learning/go-learning/gin-practise/cmd/options"
	"github.com/darianJmy/learning/go-learning/gin-practise/pkg/core"
)

var CoreV1 core.CoreV1Interface

// Setup 完成核心应用接口的设置
func Setup(o *options.Options) {
	CoreV1 = core.New(o.DB)
}
