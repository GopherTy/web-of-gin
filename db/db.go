package db

import (
	"web-of-gin/config"
	"web-of-gin/logger"
	"web-of-gin/model/auth"
	"web-of-gin/model/users"

	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

// Register 数据库注册器
type Register struct {
}

// Regist 实现 IRegister 接口，以注册获取初始化好的 db 对象。
func (Register) Regist() {
	cnf := config.Configure()

	// 初始化日志对象
	// 检查数据库配置内容是否为空。
	if cnf.DB.Driver == "" || cnf.DB.Source == "" {
		logger.Logger().Fatal("Please configure database dirver or source")
	}

	engine, err := xorm.NewEngine(cnf.DB.Driver, cnf.DB.Source)
	if err != nil {
		logger.Logger().Fatal(err.Error())
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
	logger.Logger().Info("Init db success")

	err = db.Ping()
	if err != nil {
		logger.Logger().Fatal(err.Error())
	}

	// 是否关闭用户管理
	if cnf.DB.UserManageDisable {
		return
	}

	// 用户表
	exists, err := db.IsTableExist(&users.User{})
	if err != nil {
		logger.Logger().Fatal(err.Error())
	}
	if !exists {
		err = db.CreateTables(&users.User{})
		if err != nil {
			logger.Logger().Fatal(err.Error())
		}
	}
	exists, err = db.IsTableExist(&users.UserInfo{})
	if err != nil {
		logger.Logger().Fatal(err.Error())
	}
	if !exists {
		err = db.CreateTables(&users.UserInfo{})
		if err != nil {
			logger.Logger().Fatal(err.Error())
		}
	}
	exists, err = db.IsTableExist(&users.UserLoginLog{})
	if err != nil {
		logger.Logger().Fatal(err.Error())
	}
	if !exists {
		err = db.CreateTables(&users.UserLoginLog{})
		if err != nil {
			logger.Logger().Fatal(err.Error())
		}
	}
	db.Sync2(&users.User{}, &users.UserInfo{}, &users.UserLoginLog{})

	// 用户权限表
	exists, err = db.IsTableExist(&auth.Role{})
	if err != nil {
		logger.Logger().Fatal(err.Error())
	}
	if !exists {
		err = db.CreateTables(&auth.Role{})
		if err != nil {
			logger.Logger().Fatal(err.Error())
		}
	}
	exists, err = db.IsTableExist(&auth.Auth{})
	if err != nil {
		logger.Logger().Fatal(err.Error())
	}
	if !exists {
		err = db.CreateTables(&auth.Auth{})
		if err != nil {
			logger.Logger().Fatal(err.Error())
		}
	}
	db.Sync2(&auth.Role{}, &auth.Auth{})
}

// Engine 获取 db 对象
func Engine() *xorm.Engine {
	return db
}
