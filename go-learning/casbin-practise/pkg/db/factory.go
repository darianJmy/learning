package db

import (
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db/user"
	"gorm.io/gorm"
)

type ShareDaoFactory interface {
	User() user.UserInterface
	Role() user.RoleInterface
	Menu() user.MenuInterface
}

type shareDaoFactory struct {
	db *gorm.DB
}

func (f *shareDaoFactory) User() user.UserInterface {
	return user.NewUser(f.db)
}

func (f *shareDaoFactory) Role() user.RoleInterface {
	return user.NewRole(f.db)
}

func (f *shareDaoFactory) Menu() user.MenuInterface {
	return user.NewMenu(f.db)
}

func (f *shareDaoFactory) UserRole() user.UserRoleInterface {
	return user.NewUserRole(f.db)
}

func (f *shareDaoFactory) RoleMenu() user.RoleMenuInterface {
	return user.NewRoleMenu(f.db)
}

func NewFactory(db *gorm.DB) ShareDaoFactory {
	return &shareDaoFactory{
		db: db,
	}
}
