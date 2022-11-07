package core

import (
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db"
)

type CoreV1Interface interface {
	User() UserInterface
	Role() RoleInterface
	Menu() MenuInterface
}

type core struct {
	factory db.ShareDaoFactory
}

func (c *core) User() UserInterface {
	return newUser(c)
}

func (c *core) Role() RoleInterface {
	return newRole(c)
}

func (c *core) Menu() MenuInterface {
	return newMenu(c)
}

func New(factory db.ShareDaoFactory) CoreV1Interface {
	return &core{factory}
}
