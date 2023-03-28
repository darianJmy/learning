package main

import (
	"github.com/emicklei/go-restful/v3"
	"io"
	"log"
	"net/http"
)

func main() {
	restful.DefaultContainer.Router(restful.CurlyRouter{})
	ws := new(restful.WebService)

	ws.Route(ws.GET("/resource:validate").To(validateHandler))
	ws.Route(ws.POST("/resource/{resourceId}:init").To(initHandler))
	ws.Route(ws.POST("/resource/{resourceId}:recycle").To(recycleHandler))

	restful.Add(ws)
	println("[go-restful] serve path tails from http://localhost:8080/basepath")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func validateHandler(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "validate resource completed")
}

func initHandler(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "init resource completed, resourceId: "+req.PathParameter("resourceId"))
}

func recycleHandler(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "recycle resource completed, resourceId: "+req.PathParameter("resourceId"))
}
