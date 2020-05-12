package config

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var config Config

// Config 全局JSON配置对象
type Config struct {
	DB   DataBase // 配置文件数据库对象
	HTTP HTTP     // HTTP 协议配置对象
}

// BasePath  获取项目的绝对路径
func (c *Config) BasePath() (basePath string) {
	// 获取项目绝对路径，读取配置文件。
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Fatalln(err)
	}
	path, err = filepath.Abs(path)
	if err != nil {
		log.Fatalln(err)
	}
	basePath = filepath.Dir(path)

	return
}

// HTTP HTTP协议配置对象
type HTTP struct {
	TLS      bool
	Address  string // 地址
	CertFile string //证书验证文件
	KeyFile  string // 证书
}

// DataBase 数据库连接对象
type DataBase struct {
	Driver string // 数据库驱动
	Source string // 连接字符串

	ShowSQL bool // 是否显示 SQL 语句

	MaxOpenConns int // 数据库连接池数量
	MaxIdleConns int // 数据库连接最大空闲数

	Cached int // 缓存大小
}

// Configure 获取单个配置对象
func Configure() *Config {
	return &config
}
