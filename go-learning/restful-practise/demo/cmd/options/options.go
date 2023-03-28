package options

import (
	"flag"
	"fmt"
	pixiuConfig "github.com/caoyingjunz/pixiulib/config"
	"github.com/darianJmy/restful-practise/demo/cmd/config"
	"github.com/darianJmy/restful-practise/demo/pkg/types"
	"github.com/darianJmy/restful-practise/demo/pkg/utils/log"
	"github.com/emicklei/go-restful/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
)

const (
	maxIdleConns = 10
	maxOpenConns = 100

	defaultConfigFile = "/etc/gin-practise/config.yaml"
)

type Options struct {
	ConfigFile      string
	ComponentConfig config.Config

	Addr        string
	WsContainer *restful.Container
	DB          *gorm.DB
}

func NewOptions() *Options {
	o := &Options{}
	AddFlags(o)
	flag.Parse()

	return o
}

func (o *Options) Config() *Options {
	configFile := o.ConfigFile
	if len(configFile) == 0 {
		configFile = os.Getenv("ConfigFile")
	}
	if len(configFile) == 0 {
		configFile = defaultConfigFile
	}

	c := pixiuConfig.New()
	c.SetConfigFile(configFile)
	c.SetConfigType("yaml")
	if err := c.Binding(&o.ComponentConfig); err != nil {
		panic(err)
	}

	return o
}

func (o *Options) Container() *Options {
	o.WsContainer = restful.NewContainer()
	return o
}

func (o *Options) Address() *Options {
	o.Addr = fmt.Sprintf(":%d", o.ComponentConfig.Default.Listen)
	return o
}

func (o *Options) Register() *Options {
	//if err := o.registerDatabase(); err != nil {
	//	panic(err)
	//}
	//
	o.registerLogClient()

	return o
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

func (o *Options) CheckTables() error {
	if !o.DB.Migrator().HasTable(&types.User{}) {
		if err := o.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			Migrator().CreateTable(&types.User{}); err != nil {
			return err
		}
	}
	return nil
}

func (o *Options) registerLogClient() {
	log.NewLogger(o.ComponentConfig.Default.LogDir)
}

func (o *Options) Run() {
	go func() {
		fmt.Println("starting at go-restful service")
		http.ListenAndServe(fmt.Sprintf(":%d", o.ComponentConfig.Default.Listen), o.WsContainer)
	}()
}

func AddFlags(options *Options) {
	flag.StringVar(&options.ConfigFile, "configFile", options.ConfigFile, "configFile Path")
}
