package core

import (
	"casbin-practise/pkg/types"
	"context"

	"casbin-practise/pkg/db"
	"casbin-practise/pkg/db/model"
)

type RoleInterface interface {
	GetRole(ctx context.Context, rid string) (*model.Role, error)
	ListRole(ctx context.Context) (*[]model.Role, error)
	CreateRole(ctx context.Context, obj *types.Role) error
	UpdateRole(ctx context.Context, rid string, obj *types.Role) error
}

type role struct {
	factory db.ShareDaoFactory
}

func newRole(core *core) *role {
	return &role{core.factory}
}

func (r *role) GetRole(ctx context.Context, rid string) (*model.Role, error) {
	obj, err := r.factory.Role().Get(ctx, rid)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (r *role) ListRole(ctx context.Context) (*[]model.Role, error) {
	obj, err := r.factory.Role().List(ctx)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (r *role) CreateRole(ctx context.Context, obj *types.Role) error {

	if _, err := r.factory.Role().Create(ctx, &model.Role{
		RoleName: obj.RoleName,
		RoleCode: obj.RoleCode,
		RoleDesc: obj.RoleDesc,
	}); err != nil {
		return err
	}

	return nil
}

func (r *role) UpdateRole(ctx context.Context, rid string, obj *types.Role) error {
	if _, err := r.factory.Role().Get(ctx, rid); err != nil {
		return err
	}

	if err := r.factory.Role().Update(ctx, rid, &model.Role{
		RoleName: obj.RoleName,
		RoleCode: obj.RoleCode,
		RoleDesc: obj.RoleDesc,
	}); err != nil {
		return err
	}

	return nil
}
