package user

import (
	"fmt"
	"github.com/darianJmy/restful-practise/demo/api/middleware"
	"github.com/emicklei/go-restful/v3"
)

func Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Filter(middleware.NCSACommonLogFormatLogger())
	ws.Path("/user").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("/{user-id}").To(GetUser))
	ws.Route(ws.GET("").To(ListUser))
	ws.Route(ws.POST("").To(CreateUser))
	ws.Route(ws.PUT("/{user-id}").To(UpdateUser))
	ws.Route(ws.DELETE("/{user-id}").To(DeleteUser))

	container.Add(ws)

	fmt.Println()

}
