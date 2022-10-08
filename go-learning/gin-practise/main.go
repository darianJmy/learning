package main

import (
	"github.com/darianJmy/learning/go-learning/gin-practise/api"
	"github.com/darianJmy/learning/go-learning/gin-practise/cmd"
	"github.com/darianJmy/learning/go-learning/gin-practise/cmd/options"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	s := options.NewOptions()

	if err := s.Complete(); err != nil {
		panic(err)
	}

	cmd.Setup(s)

	api.RegistryRoutes(s.GinEngine)
	s.Run()
}
