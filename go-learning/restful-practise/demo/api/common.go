package api

import (
	"github.com/darianJmy/restful-practise/demo/api/router/user"
	"github.com/emicklei/go-restful/v3"
)

func RegisterRoutes(wsContainer *restful.Container) {
	user.Register(wsContainer)
}
