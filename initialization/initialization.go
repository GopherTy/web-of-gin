package initialization

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"web-of-gin/config"

	"github.com/go-xorm/xorm"
)

var db *xorm.Engine // xorm 数据库对象

// Init 初始化数据库
func Init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := config.Configure()
	basePath := cfg.BasePath()

	// 读取配置文件
	b, err := ioutil.ReadFile(basePath + "/config.json")
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(b, cfg)
	if err != nil {
		log.Fatalln(err)
	}

	// 检查数据库配置内容是否为空。
	if cfg.DB.Driver == "" || cfg.DB.Source == "" {
		log.Fatalln("Please configure database dirver or source")
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

// DB gorm 数据操作对象
func DB() *xorm.Engine {
	return db
}
