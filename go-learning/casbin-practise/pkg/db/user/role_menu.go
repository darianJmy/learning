package user

import (
	"context"
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db/model"
	"gorm.io/gorm"
)

type RoleMenuInterface interface {
	Create(context context.Context, obj *model.RoleMenu) (*model.RoleMenu, error)
	Delete(context context.Context, uid int64) error
	Get(context context.Context, uid int64) (*model.RoleMenu, error)
	List(ctx context.Context) (*[]model.RoleMenu, error)
	Update(ctx context.Context, uid int64, user *model.RoleMenu) error
}

type roleMenu struct {
	db *gorm.DB
}

func NewRoleMenu(db *gorm.DB) RoleMenuInterface {
	return &roleMenu{db}
}

func (r *roleMenu) Create(context context.Context, obj *model.RoleMenu) (*model.RoleMenu, error) {
	return nil, nil
}

func (r *roleMenu) Delete(context context.Context, uid int64) error {
	return nil
}

func (r *roleMenu) Get(context context.Context, uid int64) (*model.RoleMenu, error) {
	return nil, nil
}

func (r *roleMenu) List(ctx context.Context) (*[]model.RoleMenu, error) {
	return nil, nil
}

func (r *roleMenu) Update(ctx context.Context, uid int64, user *model.RoleMenu) error {
	return nil
}
