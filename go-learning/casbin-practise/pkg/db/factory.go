package db

import (
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db/user"
	"gorm.io/gorm"
)

type ShareDaoFactory interface {
	User() user.UserInterface
}

type shareDaoFactory struct {
	db *gorm.DB
}

func (f *shareDaoFactory) User() user.UserInterface {
	return user.NewUser(f.db)
}

func NewFactory(db *gorm.DB) ShareDaoFactory {
	return &shareDaoFactory{
		db: db,
	}
}
