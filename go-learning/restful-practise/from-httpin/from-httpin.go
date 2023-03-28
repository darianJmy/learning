package main

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/ggicci/httpin"
	"net/http"
)

type ListUsersInput struct {
	Gender   string `in:"query=gender"`
	AgeRange []int  `in:"query=age_range"`
	IsMember bool   `in:"query=is_member"`
	Token    string `in:"header=x-client-token;query=access_token"`
}

func handleListUsers(req *restful.Request, resp *restful.Response) {
	input := req.Request.Context().Value(httpin.Input).(*ListUsersInput)
	resp.WriteAsJson(input)
}

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/users").Filter(
		restful.HttpMiddlewareHandlerToFilter(httpin.NewInput(ListUsersInput{})),
	).To(handleListUsers))

	restful.Add(ws)

	http.ListenAndServe(":8080", nil)
}
