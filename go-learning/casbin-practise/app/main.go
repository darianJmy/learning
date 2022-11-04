package main

import (
	"github.com/darianJmy/learning/go-learning/casbin-practise/api"
	cmd2 "github.com/darianJmy/learning/go-learning/casbin-practise/app/cmd"
	options2 "github.com/darianJmy/learning/go-learning/casbin-practise/app/cmd/options"
)

func main() {
	o := options2.NewOptions()

	if err := o.Complete(); err != nil {
		panic(err)
	}

	cmd2.Setup(o)

	api.RegistryRoutes(o.GinEngine)
	o.Run()
}
