package role

import (
	"github.com/emicklei/go-restful/v3"
)

type roleRouter struct {
}

func NewRouter(container *restful.Container) {
	s := &roleRouter{}
	s.initRoutes(container)
}

func (r *roleRouter) initRoutes(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/role").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("/{role-id}").To(r.GetRole))
	ws.Route(ws.GET("").To(r.ListRole))
	ws.Route(ws.POST("").To(r.CreateRole))
	ws.Route(ws.PUT("/{role-id}").To(r.UpdateRole))
	container.Add(ws)
}
