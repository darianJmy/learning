package core

import (
	"github.com/google/uuid"
	"time"

	"gorm.io/gorm"

	"github.com/darianJmy/learning/go-learning/gin-practise/pkg/types"
)

type Mysql struct {
	DB *gorm.DB
}

type MysqlGetter interface {
	Mysql() MysqlV1Interface
}

type MysqlV1Interface interface {
	CreateUser(user *types.User) error
	DeleteUser(userName string) error
	GetUser(userName string) (*types.User, error)
	ListUser() (*[]types.User, error)
	UpdateUser(userName string, user *types.User) error
}

func newMysql(c core) MysqlV1Interface {
	return &Mysql{
		DB: c.DB,
	}
}

func (m *Mysql) CreateUser(user *types.User) error {
	user.CreateTime = time.Now()
	user.UerID = uuid.New().ID()

	return m.DB.Create(&user).Error
}

func (m *Mysql) DeleteUser(userName string) error {
	var user types.User
	return m.DB.Where("username = ?", userName).Delete(&user).Error
}

func (m *Mysql) GetUser(userName string) (*types.User, error) {
	var user types.User
	if err := m.DB.Where("username = ?", userName).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *Mysql) ListUser() (*[]types.User, error) {
	var users []types.User
	if err := m.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (m *Mysql) UpdateUser(userName string, user *types.User) error {
	return m.DB.Model(&types.User{}).Where("username = ?", userName).Updates(&user).Error
}
