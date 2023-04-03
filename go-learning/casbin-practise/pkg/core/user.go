package core

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"casbin-practise/pkg/db"
	"casbin-practise/pkg/db/model"
	"casbin-practise/pkg/types"
)

type UserInterface interface {
	GetUser(ctx context.Context, uid string) (*model.User, error)
	ListUser(ctx context.Context) (*[]model.User, error)
	CreateUser(ctx context.Context, obj *types.User) error
	UpdateUser(ctx context.Context, uid string, obj *types.User) error
	DeleteUser(ctx context.Context, uid string) error
}

type user struct {
	factory db.ShareDaoFactory
}

func newUser(core *core) *user {
	return &user{core.factory}
}

func (u *user) GetUser(ctx context.Context, uid string) (*model.User, error) {
	obj, err := u.factory.User().Get(ctx, uid)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (u *user) ListUser(ctx context.Context) (*[]model.User, error) {
	obj, err := u.factory.User().List(ctx)
	if err != nil {
		return nil, err
	}
	return obj, nil
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

func (u *user) UpdateUser(ctx context.Context, uid string, obj *types.User) error {
	if _, err := u.factory.User().Get(ctx, uid); err != nil {
		return err
	}

	if obj.Password != "" || len(obj.Password) != 0 {
		// 对密码进行加密存储
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(obj.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		if err = u.factory.User().Update(ctx, uid, &model.User{
			UserName:   obj.UserName,
			UserNameCn: obj.UserNameCn,
			NickName:   obj.NickName,
			Password:   string(encryptedPassword),
			Phone:      obj.Phone,
			Email:      obj.Email,
		}); err != nil {
			return err
		}

		return nil
	}

	if err := u.factory.User().Update(ctx, uid, &model.User{
		UserName:   obj.UserName,
		UserNameCn: obj.UserNameCn,
		NickName:   obj.NickName,
		Phone:      obj.Phone,
		Email:      obj.Email,
	}); err != nil {
		return err
	}

	return nil
}

func (u *user) DeleteUser(ctx context.Context, uid string) error {

	if _, err := u.factory.User().Get(ctx, uid); err != nil {
		return err
	}

	if err := u.factory.User().Delete(ctx, uid); err != nil {
		return err
	}

	return nil
}
