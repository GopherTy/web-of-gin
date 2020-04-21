package initialization

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"web-of-gin/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB // gorm 数据库对象

// Init 初始化数据库
func Init() {
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

	// 格式正确解析，检查内容是否为空。
	if cfg.DataBase.Driver == "" || cfg.DataBase.Source == "" {
		log.Fatalln("Please configure database dirver or source")
	}

	db, err = gorm.Open(cfg.DataBase.Driver, cfg.DataBase.Source)
	if err != nil {
		log.Fatalln(err)
	}

	// 是否开启 SQL 日志
	if cfg.DataBase.ShowSQL {
		db.LogMode(true)
	}
}

// DB gorm 数据操作对象
func DB() *gorm.DB {
	return db
}
