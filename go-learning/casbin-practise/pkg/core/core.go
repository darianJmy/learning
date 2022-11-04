package core

import (
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db"
)

type CoreV1Interface interface {
	User() UserInterface
}

type core struct {
	factory db.ShareDaoFactory
}

func (c *core) User() UserInterface {
	return newUser(c)
}

func New(factory db.ShareDaoFactory) CoreV1Interface {
	return &core{factory}
}
