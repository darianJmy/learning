package core

import (
	"context"
	"fmt"
	"github.com/darianJmy/learning/go-learning/casbin-practise/api/types"
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db"
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db/model"
	"golang.org/x/crypto/bcrypt"
)

type RoleInterface interface {
	CreateRole(ctx context.Context, obj *types.Role) error
	DeleteRole(ctx context.Context, uid int64) error
	GetRole(ctx context.Context, uid int64) (*types.Role, error)
	ListRole(ctx context.Context) (*[]model.Role, error)
	UpdateRole(ctx context.Context, obj *types.Role) error
}

type role struct {
	factory db.ShareDaoFactory
}

func newRole(c *core) *role {
	return &role{c.factory}
}

func (r *role) CreateRole(ctx context.Context, obj *types.Role) error {
	if len(obj.Name) == 0 || len(obj.Password) == 0 {
		return fmt.Errorf("user name or password could not be empty")
	}

	if _, err := r.factory.Role().Create(ctx, &model.Role{
		RoleName: obj.Name,
		RoleCode: obj.UserNameCn,
		RoleDesc: obj.Nick,
	}); err != nil {
		return err
	}

	return nil
}

func (r *role) DeleteRole(ctx context.Context, uid int64) error {
	if err := r.factory.Role().Delete(ctx, uid); err != nil {
		return err
	}

	return nil
}

func (r *role) GetRole(ctx context.Context, uid int64) (*types.Role, error) {
	modelUser, err := r.factory.Role().Get(ctx, uid)
	if err != nil {
		return nil, err
	}

	return model2Type(modelUser), nil
}

func (r *role) ListRole(ctx context.Context) (*[]model.Role, error) {
	res, err := r.factory.Role().List(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *role) UpdateRole(ctx context.Context, obj *types.Role) error {
	oldUser, err := r.factory.Role().Get(ctx, obj.Id)
	if err != nil {
		return err
	}

	updates := r.parseUserUpdates(oldUser, obj)
	if len(updates) == 0 {
		return nil
	}

	if err = r.factory.Role().Update(ctx, obj.Id, updates); err != nil {
		return err
	}

	return nil
}

func model2Type(u *model.User) *types.User {
	return &types.User{
		Id:    u.Id,
		Name:  u.UserName,
		Nick:  u.NickName,
		Email: u.Email,
		Phone: u.Phone,
	}
}

func (r *role) parseUserUpdates(oldObj *model.Role, newObj *types.Role) map[string]interface{} {
	updates := make(map[string]interface{})

	if oldObj.Email != newObj.Email {
		updates["email"] = newObj.Email
	}
	if oldObj.Phone != newObj.Phone {
		updates["phone"] = newObj.Phone
	}
	return updates
}
