package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"time"
)

type User struct {
	Id, Name string
}
type UserList struct {
	Users []User
}

func main() {
	restful.Filter(globalLogging)
	restful.Add(NewUserService())
	http.ListenAndServe(":8080", nil)

}

func NewUserService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Filter(webServiceLogging).Filter(measureTime)
	ws.Route(ws.GET("/").Filter(NewCountFilter().routeCounter).To(getAllUsers))
	ws.Route(ws.GET("/{user-id}").Filter(routeLogging).Filter(NewCountFilter().routeCounter).To(findUser))
	return ws
}

//global filter
func globalLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[global-filter (logger)] %s,%s\n", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}

//webservice filter
func webServiceLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[webservice-filter (logger)] %s,%s\n", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}

func measureTime(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	now := time.Now()
	chain.ProcessFilter(req, resp)
	log.Printf("[webservice-filter (timer)] %v\n", time.Now().Sub(now))
}

func routeLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[route-filter (logger)] %s,%s\n", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}

type CountFilter struct {
	count   int
	counter chan int
}

func NewCountFilter() *CountFilter {
	c := new(CountFilter)
	c.counter = make(chan int)
	go func() {
		for {
			c.count += <-c.counter
		}
	}()
	return c
}

func (c *CountFilter) routeCounter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	c.counter <- 1
	log.Printf("[route-filter (counter)] count:%d", c.count)
	chain.ProcessFilter(req, resp)
}

func getAllUsers(req *restful.Request, resp *restful.Response) {
	log.Printf("getAllUsers")
	resp.WriteEntity(UserList{[]User{{"42", "Gandalf"}, {"3.14", "Pi"}}})
}

func findUser(req *restful.Request, resp *restful.Response) {
	log.Printf("findUser")
	resp.WriteEntity(User{"42", "Gandalf"})
}
