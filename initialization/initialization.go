package initialization

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"web-of-gin/config"

	"go.uber.org/zap"

	"github.com/go-xorm/xorm"
)

var (
	logger *zap.Logger  // zap包中的日志对象
	db     *xorm.Engine // xorm 数据库对象
)

// InitObj 用于获取初始化函数完成后单一对象
type InitObj struct {
}

// Init 初始化数据库对象和日志对象
func Init() {
	cfg := config.Configure()
	basePath := cfg.BasePath()

	// 读取配置文件
	b, err := ioutil.ReadFile(basePath + "/config.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(b, cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 初始化日志对象

	// 检查数据库配置内容是否为空。
	if cfg.DB.Driver == "" || cfg.DB.Source == "" {
		fmt.Println("Please configure database dirver or source")
		os.Exit(1)
	}

	db, err = xorm.NewEngine(cfg.DB.Driver, cfg.DB.Source)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	// 数据库相关操作

	// 设置数据库最大连接数和空闲数
	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)

	// 是否开启 SQL 日志
	if cfg.DB.ShowSQL {
		db.ShowSQL(true)
	}

	if cfg.DB.Cached != 0 {
		cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), cfg.DB.Cached)
		db.SetDefaultCacher(cacher)
	}
}

// DB xorm 数据操作对象
func (InitObj) DB() *xorm.Engine {
	return db
}

// Logger zap第三方库日志对象
func (InitObj) Logger() *zap.Logger {
	return logger
}
