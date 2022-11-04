package user

import (
	"context"
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db/model"
	"gorm.io/gorm"
	"time"
)

type MenuInterface interface {
	Create(context context.Context, obj *model.Menu) (*model.Menu, error)
	Delete(context context.Context, uid int64) error
	Get(context context.Context, uid int64) (*model.Menu, error)
	List(ctx context.Context) (*[]model.Menu, error)
	Update(ctx context.Context, uid int64, user *model.Menu) error
}

type menu struct {
	db *gorm.DB
}

func NewMenu(db *gorm.DB) MenuInterface {
	return &menu{db}
}

func (m *menu) Create(context context.Context, obj *model.Menu) (*model.Menu, error) {
	now := time.Now()
	obj.CreateTime = now

	if err := m.db.Create(obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

func (m *menu) Delete(context context.Context, uid int64) error {
	return m.db.
		Where("id = ?", uid).
		Delete(&model.Menu{}).Error
}

func (m *menu) Get(context context.Context, uid int64) (*model.Menu, error) {
	var obj model.Menu
	if err := m.db.Where("id = ?", uid).Find(&obj).Error; err != nil {
		return nil, err
	}
	return &obj, nil
}

func (m *menu) List(ctx context.Context) (*[]model.Menu, error) {
	var ms []model.Menu
	if err := m.db.Find(&ms).Error; err != nil {
		return nil, err
	}
	return &ms, nil
}

func (m *menu) Update(ctx context.Context, uid int64, menu *model.Menu) error {
	return m.db.Model(&model.Menu{}).Where("id = ?", uid).Updates(&menu).Error
}
