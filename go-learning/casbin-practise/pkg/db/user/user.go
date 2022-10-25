package user

import (
	"context"
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db/model"
	"github.com/darianJmy/learning/go-learning/gin-practise/pkg/types"
	"gorm.io/gorm"
	"time"
)

type UserInterface interface {
	Create(context context.Context, obj *model.User) (*model.User, error)
	Delete(context context.Context, uid int64) error
	Get(context context.Context, uid int64) (*model.User, error)
	List(ctx context.Context) (*[]model.User, error)
	Update(ctx context.Context, uid int64, user *model.User) error
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserInterface {
	return &user{db}
}

func (u *user) Create(context context.Context, obj *model.User) (*model.User, error) {
	now := time.Now()
	obj.CreateTime = now

	if err := u.db.Create(obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

func (u *user) Delete(context context.Context, uid int64) error {
	return u.db.
		Where("user_id = ?", uid).
		Delete(&model.User{}).Error
}

func (u *user) Get(context context.Context, uid int64) (*model.User, error) {
	var obj model.User
	if err := u.db.Where("user_id = ?", uid).Find(&obj).Error; err != nil {
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

func (u *user) Update(ctx context.Context, uid int64, user *model.User) error {
	return u.db.Model(&types.User{}).Where("user_id = ?", uid).Updates(&user).Error
}

func CheckTables(db *gorm.DB) error {
	if !db.Migrator().HasTable(&model.User{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			Migrator().CreateTable(&model.User{}); err != nil {
			return err
		}
	}
	return nil
}
