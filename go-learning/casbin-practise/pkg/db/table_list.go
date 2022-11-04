package db

import (
	"github.com/casbin/casbin/v2"
	csmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/darianJmy/learning/go-learning/casbin-practise/pkg/db/model"
	"gorm.io/gorm"
)

type tableList []error

func GetDefaultModelList(db *gorm.DB) tableList {
	return tableList{
		CheckUserTables(db),
		CheckRoleTables(db),
		CheckMenuTables(db),
		CheckUserRoleTables(db),
		CheckRoleMenuTables(db),
		CheckRbacTables(db),
	}
}

func CheckUserTables(db *gorm.DB) error {
	if !db.Migrator().HasTable(&model.User{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			Migrator().CreateTable(&model.User{}); err != nil {
			return err
		}
	}
	return nil
}

func CheckRoleTables(db *gorm.DB) error {
	if !db.Migrator().HasTable(&model.Role{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			Migrator().CreateTable(&model.Role{}); err != nil {
			return err
		}
	}
	return nil
}

func CheckMenuTables(db *gorm.DB) error {
	if !db.Migrator().HasTable(&model.Menu{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			Migrator().CreateTable(&model.Menu{}); err != nil {
			return err
		}
	}
	return nil
}

func CheckUserRoleTables(db *gorm.DB) error {
	if !db.Migrator().HasTable(&model.UserRole{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			Migrator().CreateTable(&model.UserRole{}); err != nil {
			return err
		}
	}
	return nil
}

func CheckRoleMenuTables(db *gorm.DB) error {
	if !db.Migrator().HasTable(&model.RoleMenu{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			Migrator().CreateTable(&model.RoleMenu{}); err != nil {
			return err
		}
	}
	return nil
}

func CheckRbacTables(db *gorm.DB) error {
	return InitPolicyEnforcer(db)
}

func InitPolicyEnforcer(db *gorm.DB) error {
	rbacRules :=
		`
	[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act
	
	[role_definition]
	g = _, _

	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act) || r.sub == "21220821"
	`
	// 加载鉴权规则
	m, err := csmodel.NewModelFromString(rbacRules)
	if err != nil {
		return err
	}
	// 调用gorm创建casbin_rule表
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(db, &model.Rule{}, "rules")
	if err != nil {
		return err
	}
	// 创建鉴权器enforcer（使用gorm adapter）
	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		return err
	}
	// 	加载权限
	err = enforcer.LoadPolicy()
	if err != nil {
		return err
	}

	return nil
}
