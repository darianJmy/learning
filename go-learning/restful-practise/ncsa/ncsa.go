package main

import (
	"github.com/emicklei/go-restful/v3"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var logger *log.Logger = log.New(os.Stdout, "", 0)

func NCSACommonLogFormatLogger() restful.FilterFunction {
	return func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		var username = "-"
		if req.Request.URL.User != nil {
			if name := req.Request.URL.User.Username(); name != "" {
				username = name
			}
		}
		chain.ProcessFilter(req, resp)

		logger.Printf("%s - %s [%s] \"%s %s %s\" %d %d",
			strings.Split(req.Request.RemoteAddr, ":")[0],
			username,
			time.Now().Format("02/Jan/2006:15:04:05 -0700"),
			req.Request.Method,
			req.Request.URL.RequestURI(),
			req.Request.Proto,
			resp.StatusCode(),
			resp.ContentLength(),
		)
	}
}
func main() {
	ws := new(restful.WebService)
	ws.Filter(NCSACommonLogFormatLogger())
	ws.Route(ws.GET("/ping").To(hello))
	restful.Add(ws)
	http.ListenAndServe(":8080", nil)
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "pong")
}
