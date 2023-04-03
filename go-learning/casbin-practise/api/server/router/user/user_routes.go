package user

import (
	"github.com/emicklei/go-restful/v3"
)

type userRouter struct {
}

func NewRouter(container *restful.Container) {
	s := &userRouter{}
	s.initRoutes(container)
}

func (u *userRouter) initRoutes(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/user").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("/{user-id}").To(u.GetUser))
	ws.Route(ws.GET("").To(u.ListUser))
	ws.Route(ws.POST("").To(u.CreateUser))
	ws.Route(ws.PUT("/{user-id}").To(u.UpdateUser))
	ws.Route(ws.DELETE("/{user-id}").To(u.DeleteUser))

	container.Add(ws)
}
