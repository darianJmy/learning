package user

import (
	"casbin-practise/pkg/db/model"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserInterface interface {
	Create(context context.Context, user *model.User) (*model.User, error)
	Delete(context context.Context, uid string) error
	Get(context context.Context, uid string) (*model.User, error)
	List(ctx context.Context) (*[]model.User, error)
	Update(ctx context.Context, uid string, user *model.User) error
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserInterface {
	return &user{db}
}

func (u *user) Create(context context.Context, user *model.User) (*model.User, error) {
	now := time.Now()
	user.CreateTime = now
	user.UpdateTime = now
	user.Id = uuid.New().String()

	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *user) Delete(context context.Context, uid string) error {
	return u.db.
		Where("id = ?", uid).
		Delete(&model.User{}).Error
}

func (u *user) Get(context context.Context, uid string) (*model.User, error) {
	var obj model.User
	if err := u.db.Where("id = ?", uid).Find(&obj).Error; err != nil {
		return nil, err
	}
	return &obj, nil
}

func (u *user) List(ctx context.Context) (*[]model.User, error) {
	var us []model.User
	if err := u.db.Find(&us).Error; err != nil {
		return nil, err
	}
	return &us, nil
}

func (u *user) Update(ctx context.Context, uid string, user *model.User) error {
	now := time.Now()
	user.UpdateTime = now
	user.UpdatePasswordTime = now
	return u.db.Model(&model.User{}).Where("id = ?", uid).Updates(&user).Error
}
