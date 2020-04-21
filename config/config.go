package config

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Config 全局JSON配置对象
type Config struct {
	DataBase DataBase // 配置文件数据库对象
}

// DataBase 数据库对象
type DataBase struct {
	Driver string // 数据库驱动
	Source string // 连接字符串

	ShowSQL bool // 是否显示 SQL 语句

	MaxOpenConns int // 数据库连接池数量
	MaxIdleConns int // 数据库连接最大空闲数

	Cached int // 缓存大小
}

var config Config

// Configure 获取单个配置对象
func Configure() *Config {
	return &config
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
