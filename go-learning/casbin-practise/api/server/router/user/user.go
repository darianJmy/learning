package user

import (
	"casbin-practise/pkg/global"
	"context"
	"github.com/emicklei/go-restful/v3"

	"casbin-practise/pkg/types"
)

func (u *userRouter) GetUser(req *restful.Request, resp *restful.Response) {
	uid := req.PathParameter("user-id")
	global.Corev1.User().GetUser(context.TODO(), uid)

	resp.WriteHeaderAndJson(200, uid, "application/json")
}

func (u *userRouter) ListUser(req *restful.Request, resp *restful.Response) {
	users, err := global.Corev1.User().ListUser(context.TODO())
	if err != nil {
		resp.WriteHeaderAndEntity(400, nil)
		return
	}

	resp.WriteHeaderAndJson(200, users, "application/json")
}

func (u *userRouter) CreateUser(req *restful.Request, resp *restful.Response) {
	var user types.User

	if err := req.ReadEntity(&user); err != nil {
		resp.WriteHeaderAndEntity(400, nil)
		return
	}

	if err := global.Corev1.User().CreateUser(context.TODO(), &user); err != nil {
		resp.WriteHeaderAndEntity(400, nil)
		return
	}

	resp.WriteHeaderAndEntity(200, nil)
}

func (u *userRouter) UpdateUser(req *restful.Request, resp *restful.Response) {
	var user types.User
	if err := req.ReadEntity(&user); err != nil {
		resp.WriteHeaderAndEntity(400, nil)
		return
	}

	uid := req.PathParameter("user-id")
	global.Corev1.User().UpdateUser(context.TODO(), uid, &user)
}

func (u *userRouter) DeleteUser(req *restful.Request, resp *restful.Response) {
	uid := req.PathParameter("user-id")
	if err := global.Corev1.User().DeleteUser(context.TODO(), uid); err != nil {
		resp.WriteHeaderAndEntity(400, nil)
		return
	}

	resp.WriteHeaderAndEntity(200, nil)
}
