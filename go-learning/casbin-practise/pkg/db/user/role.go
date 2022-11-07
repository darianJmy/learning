package user

import (
	"context"
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db/model"
	"gorm.io/gorm"
	"time"
)

type RoleInterface interface {
	Create(context context.Context, obj *model.Role) (*model.Role, error)
	Delete(context context.Context, uid int64) error
	Get(context context.Context, uid int64) (*model.Role, error)
	List(ctx context.Context) (*[]model.Role, error)
	Update(ctx context.Context, uid int64, updates map[string]interface{}) error
}

type role struct {
	db *gorm.DB
}

func NewRole(db *gorm.DB) RoleInterface {
	return &role{db}
}

func (r *role) Create(context context.Context, obj *model.Role) (*model.Role, error) {
	now := time.Now()
	obj.CreateTime = now

	if err := r.db.Create(obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *role) Delete(context context.Context, uid int64) error {
	return r.db.
		Where("id = ?", uid).
		Delete(&model.Role{}).Error
}

func (r *role) Get(context context.Context, uid int64) (*model.Role, error) {
	var obj model.Role
	if err := r.db.Where("id = ?", uid).Find(&obj).Error; err != nil {
		return nil, err
	}
	return &obj, nil
}

func (r *role) List(ctx context.Context) (*[]model.Role, error) {
	var rs []model.Role
	if err := r.db.Find(&rs).Error; err != nil {
		return nil, err
	}
	return &rs, nil
}

func (r *role) Update(ctx context.Context, uid int64, updates map[string]interface{}) error {
	return r.db.Model(&model.Role{}).Where("id = ?", uid).Updates(updates).Error
}
