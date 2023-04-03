package api

import (
	"casbin-practise/api/server/middleware"
	"casbin-practise/api/server/router/role"
	"github.com/emicklei/go-restful/v3"

	"casbin-practise/api/server/router/user"
)

func RegisterRoutes(wsContainer *restful.Container) {

	wsContainer.Filter(middleware.NCSACommonLogFormatLogger())
	user.NewRouter(wsContainer)
	role.NewRouter(wsContainer)
}
