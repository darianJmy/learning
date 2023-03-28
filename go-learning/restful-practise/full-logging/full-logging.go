package main

import (
	"bytes"
	"github.com/emicklei/go-restful/v3"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Id   int
	Name string
}

func main() {
	restful.Add(newUserService())
	http.ListenAndServe(":8080", nil)

}

func newUserService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.POST("/").Filter(bodyLogFilter).To(createUser))
	return ws
}

func bodyLogFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	inBody, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		resp.WriteError(http.StatusBadRequest, err)
	}

	req.Request.Body = ioutil.NopCloser(bytes.NewReader(inBody))

	c := NewResponseCapture(resp.ResponseWriter)

	resp.ResponseWriter = c

	chain.ProcessFilter(req, resp)

	log.Println("Request body:", string(inBody))
	log.Println("Response body:", string(c.Bytes()))
}

func createUser(req *restful.Request, resp *restful.Response) {
	u := new(User)
	err := req.ReadEntity(u)
	log.Println("createUser", err, u)
	resp.WriteEntity(u)
}

type ResponseCapture struct {
	http.ResponseWriter
	wroteHeader bool
	status      int
	body        *bytes.Buffer
}

func NewResponseCapture(w http.ResponseWriter) *ResponseCapture {
	return &ResponseCapture{
		ResponseWriter: w,
		wroteHeader:    false,
		body:           new(bytes.Buffer),
	}
}

func (c *ResponseCapture) Header() http.Header {
	return c.ResponseWriter.Header()
}

func (c *ResponseCapture) Write(data []byte) (int, error) {
	if !c.wroteHeader {
		c.WriteHeader(http.StatusOK)
	}
	c.body.Write(data)
	return c.ResponseWriter.Write(data)
}

func (c *ResponseCapture) WriteHeader(statusCode int) {
	c.status = statusCode
	c.wroteHeader = true
	c.ResponseWriter.WriteHeader(statusCode)
}

func (c *ResponseCapture) Bytes() []byte {
	return c.body.Bytes()
}

func (c *ResponseCapture) StatusCode() int {
	return c.status
}
