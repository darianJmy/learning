package types

import "time"

type UserUri struct {
	UserName string `uri:"username" binding:"required"`
}

type User struct {
	UerID              uint32    `json:"user_id" gorm:"column:user_id;type:int;primary_key;auto_increment;comment:用户ID"`
	UserName           string    `json:"username" gorm:"column:username;type:varchar(255);default:null;comment:英文名"`
	UserNameCn         string    `json:"username_cn" gorm:"column:username_cn;type:varchar(255);default:null;comment:中文名"`
	NickName           string    `json:"nickname" gorm:"column:nickname;type:varchar(255);default:null;comment:昵称"`
	Password           string    `json:"password" gorm:"column:password;type:varchar(255);default:null;comment:密码"`
	Salt               string    `json:"salt" gorm:"column:salt;type:varchar(255);default:null;comment:随机盐"`
	Phone              string    `json:"phone" gorm:"column:phone;type:varchar(255);default:null;comment:手机"`
	Email              string    `json:"email" gorm:"column:email;type:varchar(255);default:null;comment:邮箱"`
	CreateTime         time.Time `json:"create_time" gorm:"column:create_time;type:datetime;default:current_timestamp;comment:创建时间"`
	UpdateTime         time.Time `json:"update_time" gorm:"column:update_time;type:datetime;default:null on update current_timestamp;comment:修改时间"`
	UpdatePasswordTime time.Time `json:"update_password_time" gorm:"column:update_password_time;type:datetime;default:null;comment:更新密码时间"`
	LockFlag           int       `json:"lock_flag" gorm:"column:lock_flag;default:0;comment:是否锁定 0:正常 9:锁定 2:冻结"`
}
