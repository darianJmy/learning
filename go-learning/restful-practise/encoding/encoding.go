package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"net/http"
)

type User struct {
	Id, Name string
}

type UserList struct {
	Users []User
}

func main() {
	restful.Add(NewUserService())
	http.ListenAndServe(":8080", nil)
}

func NewUserService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/user").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("/{user-id}").Filter(encodingFilter).To(findUser))

	return ws
}

func encodingFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	fmt.Printf("[encoding-filter] %s,%s\n", req.Request.Method, req.Request.URL)

	compress, _ := restful.NewCompressingResponseWriter(resp.ResponseWriter, restful.ENCODING_GZIP)
	resp.ResponseWriter = compress
	defer func() {
		compress.Close()
	}()
	chain.ProcessFilter(req, resp)
}

func findUser(req *restful.Request, resp *restful.Response) {
	fmt.Print("findUser")
	resp.WriteEntity(User{"42", "Gandalf"})
}
