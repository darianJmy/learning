package main

import (
	"github.com/emicklei/go-restful/v3"
	"io"
	"net/http"
)

func main() {
	ws := new(restful.WebService)
	ws.Filter(restful.NoBrowserCacheFilter)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)
	http.ListenAndServe(":8080", nil)
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}
