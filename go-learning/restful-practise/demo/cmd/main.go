package main

import (
	"github.com/darianJmy/restful-practise/demo/api"
	"github.com/darianJmy/restful-practise/demo/cmd/options"
)

func main() {
	s := options.NewOptions()
	s.Config().
		Container().
		Address().
		Register()

	api.RegisterRoutes(s.WsContainer)

	s.Run()

	select {}
}
