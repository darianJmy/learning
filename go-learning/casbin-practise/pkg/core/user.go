package core

import (
	"context"
	"github.com/darianJmy/learning/go-learning/casbin-practise/api/types"
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db"
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db/model"
	"golang.org/x/crypto/bcrypt"
)

type UserInterface interface {
	CreateUser(ctx context.Context, obj *types.User) error
	DeleteUser(ctx context.Context, uid int64) error
	GetUser(ctx context.Context, uid int64) (*model.User, error)
}

type user struct {
	factory db.ShareDaoFactory
}

func newUser(core *core) *user {
	return &user{core.factory}
}

func (u *user) CreateUser(ctx context.Context, obj *types.User) error {

	// 对密码进行加密存储
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(obj.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = u.factory.User().Create(ctx, &model.User{
		UserName:   obj.UserName,
		UserNameCn: obj.UserNameCn,
		NickName:   obj.NickName,
		Password:   string(encryptedPassword),
		Phone:      obj.Phone,
		Email:      obj.Email,
	})

	if err != nil {
		return err
	}

	return nil
}

func (u *user) DeleteUser(ctx context.Context, uid int64) error {

	if _, err := u.factory.User().Get(ctx, uid); err != nil {
		return err
	}

	if err := u.factory.User().Delete(ctx, uid); err != nil {
		return err
	}

	return nil
}

func (u *user) GetUser(ctx context.Context, uid int64) (*model.User, error) {
	obj, err := u.factory.User().Get(ctx, uid)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
