package user

import (
	"github.com/darianJmy/restful-practise/demo/pkg/types"
	"github.com/emicklei/go-restful/v3"
	"io"
)

func GetUser(req *restful.Request, resp *restful.Response) {
	userId := req.PathParameter("user-id")

	io.WriteString(resp, userId)
}

func ListUser(req *restful.Request, resp *restful.Response) {
	resp.WriteHeaderAndEntity(200, types.User{Name: "jixingxing", Age: 18})
}

func CreateUser(req *restful.Request, resp *restful.Response) {
	var user types.User
	if err := req.ReadEntity(&user); err != nil {
		resp.WriteHeaderAndEntity(400, nil)
		return
	}

	resp.WriteHeaderAndEntity(200, user)
}

func UpdateUser(req *restful.Request, resp *restful.Response) {
	var user types.User
	userId := req.PathParameter("user-id")
	if err := req.ReadEntity(&user); err != nil {
		resp.WriteHeaderAndJson(400, nil, "application/json")
		return
	}

	user.Id = userId

	resp.WriteHeaderAndJson(200, user, "application/json")
}

func DeleteUser(req *restful.Request, resp *restful.Response) {
	resp.WriteHeaderAndEntity(200, nil)
}
