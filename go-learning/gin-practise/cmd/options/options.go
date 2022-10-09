package options

import (
	"fmt"
	"os"

	pixiuConfig "github.com/caoyingjunz/pixiulib/config"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/darianJmy/learning/go-learning/gin-practise/cmd/config"
	"github.com/darianJmy/learning/go-learning/gin-practise/pkg/log"
	"github.com/darianJmy/learning/go-learning/gin-practise/pkg/types"
)

const (
	maxIdleConns = 10
	maxOpenConns = 100

	defaultConfigFile = "/etc/gin-practise/config.yaml"
)

type Options struct {
	ConfigFile      string
	ComponentConfig config.Config

	GinEngine *gin.Engine

	DB *gorm.DB
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) Complete() error {
	configFile := o.ConfigFile
	if len(configFile) == 0 {
		configFile = os.Getenv("ConfigFile")
	}
	if len(configFile) == 0 {
		configFile = defaultConfigFile
	}

	// 解析 yaml 文件
	c := pixiuConfig.New()
	c.SetConfigFile(configFile)
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
	if err := o.registerDatabase(); err != nil {
		return err
	}
	if err := o.registerGinEngine(); err != nil {
		return err
	}
	if err := o.registerLogClient(); err != nil {
		return err
	}

	return nil
}

func (o *Options) registerLogClient() error {
	return log.ConsoleLogger(o.ComponentConfig.Default.LogDir)
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

	return o.CheckTables()
}

func (o *Options) registerGinEngine() error {
	o.GinEngine = gin.Default()

	return nil
}

func (o *Options) Run() {
	o.GinEngine.Run(fmt.Sprintf(":%d", o.ComponentConfig.Default.Listen))
}

func (o *Options) CheckTables() error {
	if !o.DB.Migrator().HasTable(&types.User{}) {
		if err := o.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			Migrator().CreateTable(&types.User{}); err != nil {
			return err
		}
	}
	return nil
}
