package options

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"

	pixiuConfig "github.com/caoyingjunz/pixiulib/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"casbin-practise/cmd/app/config"
	"casbin-practise/pkg/db"
)

const (
	maxIdleConns = 10
	maxOpenConns = 100
)

type Options struct {
	ConfigFile      string
	ComponentConfig config.Config

	DB      *gorm.DB
	Factory db.ShareDaoFactory

	Container *restful.Container
}

func NewOptions(configFile string) *Options {
	return &Options{
		ConfigFile: configFile,
	}
}

func (o *Options) NewContainer(container *restful.Container) *Options {
	o.Container = container
	return o
}

func (o *Options) Complete() error {
	// 解析 yaml 文件
	c := pixiuConfig.New()
	c.SetConfigFile(o.ConfigFile)
	c.SetConfigType("yaml")
	if err := c.Binding(&o.ComponentConfig); err != nil {
		return err
	}

	// 注册依赖组件
	if err := o.register(); err != nil {
		return err
	}
	return nil
}

func (o *Options) register() error {
	return o.registerDatabase()
}

func (o *Options) registerDatabase() error {
	sqlConfig := o.ComponentConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		sqlConfig.User,
		sqlConfig.Password,
		sqlConfig.Host,
		sqlConfig.Port,
		sqlConfig.Database)

	var err error
	o.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 设置数据库连接池
	sqlDB, err := o.DB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)

	o.Factory = db.NewFactory(o.DB)

	return o.checkTables()
}

func (o *Options) checkTables() error {
	modelList := db.GetDefaultModelList(o.DB)
	for _, err := range modelList {
		if err != nil {
			return err
		}
	}
	return nil
}
