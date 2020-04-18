package initialization

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"
)

// Config 全局JSON配置对象
type Config struct {
	DataBase DataBase
}

// DataBase 数据库配置配置
type DataBase struct {
	Driver string
	Source string
}

// Manager 初始化对象
type Manager struct {
	config *Config  // 全局配置对象
	db     *gorm.DB // 数据库对象
}

var _manager Manager // 全局对象

// Single 获取单个 Manager 对象
func Single() Manager {
	return _manager
}

// Init 初始化数据库
func (m *Manager) Init() {
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalln(err)
	}

	cfg := &Config{}
	err = json.Unmarshal(b, cfg)
	if err != nil {
		log.Fatalln(err)
	}

	// 格式正确解析，检查内容是否为空。
	if cfg.DataBase.Driver == "" || cfg.DataBase.Source == "" {
		log.Fatalln("Please configure database dirver or source")
	}

	tmpDB, err := gorm.Open(cfg.DataBase.Driver, cfg.DataBase.Source)
	if err != nil {
		log.Fatalln(err)
	}
	defer tmpDB.Close()

	m.db = tmpDB
	m.config = cfg
}

// DB 获取数据库对象
func (m *Manager) DB() *gorm.DB {
	return m.db
}

// Configure 获取 json 对象
func (m *Manager) Configure() *Config {
	return m.config
}
