package global

import (
	"casbin-practise/cmd/app/options"
	"casbin-practise/pkg/core"
)

var Corev1 core.CoreV1Interface

func Setup(o *options.Options) {
	Corev1 = core.New(o.Factory)
}
