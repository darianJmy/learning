package user

import (
	"context"
	"gorm.io/gorm"

	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db/model"
)

type UserRoleInterface interface {
	Create(context context.Context, obj *model.UserRole) (*model.UserRole, error)
	Delete(context context.Context, uid int64) error
	Get(context context.Context, uid int64) (*model.UserRole, error)
	List(ctx context.Context) (*[]model.UserRole, error)
	Update(ctx context.Context, uid int64, user *model.UserRole) error
}

type userRole struct {
	db *gorm.DB
}

func NewUserRole(db *gorm.DB) UserRoleInterface {
	return &userRole{db}
}

func (u *userRole) Create(context context.Context, obj *model.UserRole) (*model.UserRole, error) {
	return nil, nil

}

func (u *userRole) Delete(context context.Context, uid int64) error {
	return nil
}

func (u *userRole) Get(context context.Context, uid int64) (*model.UserRole, error) {
	return nil, nil
}

func (u *userRole) List(ctx context.Context) (*[]model.UserRole, error) {
	return nil, nil
}

func (u *userRole) Update(ctx context.Context, uid int64, user *model.UserRole) error {
	return nil
}
