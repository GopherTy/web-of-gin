package db

import (
	"fmt"
	"os"
	"web-of-gin/model/auth"
	"web-of-gin/model/users"
	"web-of-gin/module/configs"
	"web-of-gin/module/logger"

	_ "github.com/go-sql-driver/mysql" // 模块注册导入数据库驱动
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

// Register 注册器
type Register struct {
}

// Regist 实现 IRegister 接口，以注册获取初始化好的 db 对象。
func (Register) Regist() {
	cnf := configs.Instance()

	// 初始化日志对象
	// 检查数据库配置内容是否为空。
	if cnf.DB.Driver == "" || cnf.DB.Source == "" {
		logger.Instance().Fatal("Please configure database dirver or source")
	}

	engine, err := xorm.NewEngine(cnf.DB.Driver, cnf.DB.Source)
	if err != nil {
		logger.Instance().Fatal(err.Error())
	}

	// 设置数据库最大连接数和空闲数
	engine.SetMaxOpenConns(cnf.DB.MaxOpenConns)
	engine.SetMaxIdleConns(cnf.DB.MaxIdleConns)

	// 是否开启 SQL 日志
	if cnf.DB.ShowSQL {
		engine.ShowSQL(true)
	}

	// 是否开启缓存
	if cnf.DB.Cached != 0 {
		cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), cnf.DB.Cached)
		engine.SetDefaultCacher(cacher)
	}

	db = engine
	logger.Instance().Info("Init db success")

	err = db.Ping()
	if err != nil {
		logger.Instance().Fatal(err.Error())
	}

	// 是否关闭用户管理
	if cnf.DB.UserManageDisable {
		return
	}

	// 创建用户相关的表
	err = createTable(&users.User{}, &users.UserInfo{}, &users.UserLoginLog{},
		&auth.Role{}, &auth.Auth{})
	if err != nil {
		logger.Instance().Fatal(err.Error())
	}
	// 同步表结构
	err = db.Sync2(&users.User{}, &users.UserInfo{}, &users.UserLoginLog{},
		&auth.Role{}, &auth.Auth{})
	if err != nil {
		logger.Instance().Fatal(err.Error())
	}
}

func createTable(beans ...interface{}) (err error) {
	var exists bool
	for _, bean := range beans {
		exists, err = db.IsTableExist(bean)
		if err != nil {
			return
		}
		if !exists {
			err = db.CreateTables(bean)
			if err != nil {
				return
			}
		}
	}
	return
}

// Engine 获取 db 对象
func Engine() *xorm.Engine {
	if db == nil {
		fmt.Println("Register db module failed.")
		os.Exit(1)
	}
	return db
}
