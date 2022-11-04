package model

import (
	"time"
)

// 默认格式
type Model struct {
	Id              int64     `json:"id" gorm:"column:id;type:int();primary_key;auto_increment;comment:ID"`
	CreateTime      time.Time `json:"create_time" gorm:"column:create_time;type:datetime;default:current_timestamp;comment:创建时间"`
	UpdateTime      time.Time `json:"update_time" gorm:"column:update_time;type:datetime;default:null on update current_timestamp;comment:修改时间"`
	ResourceVersion int64     `json:"resource_version" gorm:"column:resource_version;type:int;default:1;comment:资源版本"`
}

// 用户表
type User struct {
	Model
	UserName           string    `json:"username" gorm:"column:username;type:varchar(64);default:null;comment:英文名"`
	UserNameCn         string    `json:"username_cn" gorm:"column:username_cn;type:varchar(64);default:null;comment:中文名"`
	NickName           string    `json:"nickname" gorm:"column:nickname;type:varchar(64);default:null;comment:昵称"`
	Password           string    `json:"password" gorm:"column:password;type:varchar(255);default:null;comment:密码"`
	Salt               string    `json:"salt" gorm:"column:salt;type:varchar(64);default:null;comment:随机盐"`
	Phone              string    `json:"phone" gorm:"column:phone;type:varchar(64);default:null;comment:手机"`
	Email              string    `json:"email" gorm:"column:email;type:varchar(64);default:null;comment:邮箱"`
	UpdatePasswordTime time.Time `json:"update_password_time" gorm:"column:update_password_time;type:datetime;default:null;comment:更新密码时间"`
	LockFlag           int64     `json:"lock_flag" gorm:"column:lock_flag;type:int;default:0;comment:是否锁定 0:正常 9:锁定 2:冻结"`
}

// 角色表
type Role struct {
	Model
	RoleName string `gorm:"column:role_name;type:varchar(64)" json:"role_name"`
	RoleCode string `gorm:"column:role_code;type:varchar(64)" json:"role_code"`
	RoleDesc string `gorm:"column:role_desc;type:varchar(255)" json:"role_desc"`
	DsType   string `gorm:"column:ds_type;type:char(1);default:2" json:"ds_type"`
	DsScope  string `gorm:"column:ds_scope;type:varchar(255)" json:"ds_scope"`
	BuildIn  string `gorm:"column:build_in;type:char(1);default:1;comment:是否内置  0:是 1:否" json:"build_in"`
	DelFlag  string `gorm:"column:del_flag;type:char(1);default:0" json:"del_flag"`
	TenantId int64  `gorm:"column:tenant_id;type:int" json:"tenant_id"`
}

// 菜单表
type Menu struct {
	Model
	Name       string `gorm:"column:name;type:varchar(32)" json:"name"`
	Permission string `gorm:"column:permission;type:varchar(32)" json:"permission"`
	Path       string `gorm:"column:path;type:varchar(128)" json:"path"`
	ParentId   int64  `gorm:"column:parent_id;type:int(11);comment:父菜单ID" json:"parent_id"`
	Icon       string `gorm:"column:icon;type:varchar(32)" json:"icon"`
	Sort       int64  `gorm:"column:sort;type:int(11);default:1;comment:排序值" json:"sort"`
	KeepAlive  string `gorm:"column:keep_alive;type:char(1);default:0" json:"keep_alive"`
	Type       string `gorm:"column:type;type:char(1);default:0" json:"type"`
	BuildIn    string `gorm:"column:build_in;type:char(1);default:1;comment:是否内置  0:是 1:否" json:"build_in"`
	DelFlag    string `gorm:"column:del_flag;type:char(1);default:0" json:"del_flag"`
}

// 用户角色关联表
type UserRole struct {
	Model
	UserId int64 `gorm:"column:user_id;type:int;primary_key;comment:用户ID" json:"user_id"`
	RoleId int64 `gorm:"column:role_id;type:int;comment:角色ID;NOT NULL" json:"role_id"`
}

// 角色菜单关联表
type RoleMenu struct {
	Model
	RoleId int64 `gorm:"column:role_id;type:int;primary_key;comment:角色ID" json:"role_id"`
	MenuId int64 `gorm:"column:menu_id;type:int;comment:菜单ID;NOT NULL" json:"menu_id"`
}

// 权限控制表
type Rule struct {
	Model
	PType  string `json:"ptype" gorm:"column:ptype;size:100" description:"策略类型"`
	Role   string `json:"role" gorm:"column:v0;size:100" description:"角色"`
	Path   string `json:"path" gorm:"column:v1;size:100" description:"api路径"`
	Method string `json:"method" gorm:"column:v2;size:100" description:"访问方法"`
	V3     string `gorm:"column:v3;size:100"`
	V4     string `gorm:"column:v4;size:100"`
	V5     string `gorm:"column:v5;size:100"`
}
