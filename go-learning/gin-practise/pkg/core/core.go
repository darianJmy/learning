package core

import "gorm.io/gorm"

type CoreV1Interface interface {
	MysqlGetter
}

type core struct {
	DB *gorm.DB
}

func (core core) Mysql() MysqlV1Interface {
	return newMysql(core)
}

func New(DB *gorm.DB) CoreV1Interface {
	return &core{
		DB: DB,
	}
}
