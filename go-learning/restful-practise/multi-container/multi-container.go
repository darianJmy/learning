package main

import (
	"github.com/emicklei/go-restful/v3"
	"io"
	"net/http"
)

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)
	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	container2 := restful.NewContainer()
	ws2 := new(restful.WebService)
	ws2.Route(ws2.GET("/hello").To(hello2))
	container2.Add(ws)
	server := &http.Server{Addr: ":8081", Handler: container2}
	server.ListenAndServe()
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "default world")
}

func hello2(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "second world")
}
