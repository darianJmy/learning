package core

import (
	"context"
	"fmt"
	"github.com/darianJmy/learning/go-learning/casbin-practise/api/types"
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db"
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db/model"
	"golang.org/x/crypto/bcrypt"
)

type UserInterface interface {
	CreateUser(ctx context.Context, obj *types.User) error
	DeleteUser(ctx context.Context, uid int64) error
	GetUser(ctx context.Context, uid int64) (*types.User, error)
	ListUser(ctx context.Context) (*[]model.User, error)
	UpdateUser(ctx context.Context, obj *types.User) error
}

type user struct {
	factory db.ShareDaoFactory
}

func newUser(c *core) *user {
	return &user{c.factory}
}

func (u *user) CreateUser(ctx context.Context, obj *types.User) error {
	if len(obj.Name) == 0 || len(obj.Password) == 0 {
		return fmt.Errorf("user name or password could not be empty")
	}

	// 对密码进行加密存储
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(obj.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if _, err = u.factory.User().Create(ctx, &model.User{
		UserName:   obj.Name,
		UserNameCn: obj.UserNameCn,
		NickName:   obj.Nick,
		Password:   string(encryptedPassword),
		Phone:      obj.Phone,
		Email:      obj.Email,
	}); err != nil {
		return err
	}

	return nil
}

func (u *user) DeleteUser(ctx context.Context, uid int64) error {
	if err := u.factory.User().Delete(ctx, uid); err != nil {
		return err
	}

	return nil
}

func (u *user) GetUser(ctx context.Context, uid int64) (*types.User, error) {
	modelUser, err := u.factory.User().Get(ctx, uid)
	if err != nil {
		return nil, err
	}

	return model2Type(modelUser), nil
}

func (u *user) ListUser(ctx context.Context) (*[]model.User, error) {
	res, err := u.factory.User().List(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *user) UpdateUser(ctx context.Context, obj *types.User) error {
	oldUser, err := u.factory.User().Get(ctx, obj.Id)
	if err != nil {
		return err
	}

	updates := u.parseUserUpdates(oldUser, obj)
	if len(updates) == 0 {
		return nil
	}

	if err = u.factory.User().Update(ctx, obj.Id, updates); err != nil {
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

func (u *user) parseUserUpdates(oldObj *model.User, newObj *types.User) map[string]interface{} {
	updates := make(map[string]interface{})

	if oldObj.Email != newObj.Email {
		updates["email"] = newObj.Email
	}
	if oldObj.Phone != newObj.Phone {
		updates["phone"] = newObj.Phone
	}
	return updates
}
