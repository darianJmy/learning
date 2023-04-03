package middleware

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/emicklei/go-restful/v3"
)

var logger = log.New(os.Stdout, "", 0)

func NCSACommonLogFormatLogger() restful.FilterFunction {
	return func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		chain.ProcessFilter(req, resp)

		logger.Printf("%s - [%s] \"%s %s %s\" %d %d",
			strings.Split(req.Request.RemoteAddr, ":")[0],
			time.Now().Format("02/Jan/2006:15:04:05 -0700"),
			req.Request.Method,
			req.Request.URL.RequestURI(),
			req.Request.Proto,
			resp.StatusCode(),
			resp.ContentLength(),
		)
	}
}
