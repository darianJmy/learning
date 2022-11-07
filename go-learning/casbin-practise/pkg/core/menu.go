package core

import (
	"context"
	"fmt"
	"github.com/darianJmy/learning/go-learning/casbin-practise/api/types"
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db"
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db/model"
	"golang.org/x/crypto/bcrypt"
)

type MenuInterface interface {
	CreateMenu(ctx context.Context, obj *types.Menu) error
	DeleteMenu(ctx context.Context, uid int64) error
	GetMenu(ctx context.Context, uid int64) (*types.Menu, error)
	ListMenu(ctx context.Context) (*[]model.Menu, error)
	UpdateMenu(ctx context.Context, obj *types.Menu) error
}

type menu struct {
	factory db.ShareDaoFactory
}

func newMenu(c *core) *menu {
	return &menu{c.factory}
}

func (m *menu) CreateMenu(ctx context.Context, obj *types.Menu) error {
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

func (m *menu) DeleteMenu(ctx context.Context, uid int64) error {
	if err := u.factory.User().Delete(ctx, uid); err != nil {
		return err
	}

	return nil
}

func (u *menu) GetMenu(ctx context.Context, uid int64) (*types.Menu, error) {
	modelUser, err := u.factory.User().Get(ctx, uid)
	if err != nil {
		return nil, err
	}

	return model2Type(modelUser), nil
}

func (m *menu) ListMenu(ctx context.Context) (*[]model.Menu, error) {
	res, err := u.factory.User().List(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *menu) UpdateMenu(ctx context.Context, obj *types.Menu) error {
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

func (m *menu) parseUserUpdates(oldObj *model.User, newObj *types.User) map[string]interface{} {
	updates := make(map[string]interface{})

	if oldObj.Email != newObj.Email {
		updates["email"] = newObj.Email
	}
	if oldObj.Phone != newObj.Phone {
		updates["phone"] = newObj.Phone
	}
	return updates
}
