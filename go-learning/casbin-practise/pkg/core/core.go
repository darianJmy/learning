package core

import (
	"casbin-practise/pkg/db"
)

type CoreV1Interface interface {
	User() UserInterface
	Role() RoleInterface
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

func New(factory db.ShareDaoFactory) CoreV1Interface {
	return &core{factory}
}
