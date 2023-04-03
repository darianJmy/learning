package role

import (
	"casbin-practise/pkg/global"
	"casbin-practise/pkg/types"
	"context"
	"fmt"
	"github.com/emicklei/go-restful/v3"
)

func (r *roleRouter) GetRole(req *restful.Request, resp *restful.Response) {
	rid := req.PathParameter("role-id")

	role, err := global.Corev1.Role().GetRole(context.TODO(), rid)
	if err != nil {
		resp.WriteHeaderAndEntity(400, nil)
		return
	}

	resp.WriteHeaderAndJson(200, role, "application/json")
}

func (r *roleRouter) ListRole(req *restful.Request, resp *restful.Response) {
	users, err := global.Corev1.Role().ListRole(context.TODO())
	if err != nil {
		resp.WriteHeaderAndEntity(400, nil)
		return
	}

	resp.WriteHeaderAndJson(200, users, "application/json")
}

func (r *roleRouter) CreateRole(req *restful.Request, resp *restful.Response) {
	var role types.Role

	if err := req.ReadEntity(&role); err != nil {
		resp.WriteHeaderAndEntity(400, nil)
		return
	}

	if err := global.Corev1.Role().CreateRole(context.TODO(), &role); err != nil {
		resp.WriteHeaderAndEntity(400, nil)
		return
	}

	resp.WriteHeaderAndJson(200, nil, "application/json")
}

func (r *roleRouter) UpdateRole(req *restful.Request, resp *restful.Response) {
	var role types.Role
	if err := req.ReadEntity(&role); err != nil {
		resp.WriteHeaderAndEntity(400, nil)
		return
	}

	rid := req.PathParameter("role-id")
	fmt.Print(rid, "hello")
	if err := global.Corev1.Role().UpdateRole(context.TODO(), rid, &role); err != nil {
		resp.WriteHeaderAndEntity(400, nil)
		return
	}

	resp.WriteHeaderAndJson(200, nil, "application/json")
}
