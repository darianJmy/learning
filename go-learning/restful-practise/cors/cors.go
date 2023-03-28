package main

import (
	"github.com/emicklei/go-restful/v3"
	"io"
	"net/http"
)

type UserResource struct {
}

func (u *UserResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/user").
		Consumes("*/*").
		Produces("*/*")

	ws.Route(ws.GET("/{user-id}").To(u.nop))
	ws.Route(ws.POST("").To(u.nop))
	ws.Route(ws.PUT("/{user-id}").To(u.nop))
	ws.Route(ws.DELETE("/{user-id}").To(u.nop))

	container.Add(ws)
}

func (u *UserResource) nop(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp.ResponseWriter, "this would be a normal response")
}

func main() {
	wsContainer := restful.NewContainer()
	u := UserResource{}
	u.RegisterTo(wsContainer)

	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST"},
		CookiesAllowed: false,
		Container:      wsContainer,
	}
	wsContainer.Filter(cors.Filter)

	wsContainer.Filter(wsContainer.OPTIONSFilter)

	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	server.ListenAndServe()
}
