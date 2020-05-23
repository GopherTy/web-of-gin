package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"web-of-gin/config"

	"go.uber.org/zap"
)

var logger *zap.Logger

// Register 日志注册器
type Register struct {
}

// Regist 实现 IRegister 接口，以注册获取初始化好的 logger 对象。
func (Register) Regist() {
	cfg := config.Configure()

	// 日志等级 "debug", "info", "warn",
	// "error", "dpanic", "panic", and "fatal"
	jsonBytes := []byte(`
		{
			"level":"` + cfg.Logger.Level + `", 
			"encoding": "json",
			"outputPaths": ["stdout", "/tmp/logs"],
			"errorOutputPaths": ["stderr"],
			"initialFields": {"foo": "bar"},
			"encoderConfig": {
			  "messageKey": "message",
			  "levelKey": "level",
			  "levelEncoder": "lowercase"
			}
		}
	`)

	var zapCfg zap.Config
	if err := json.Unmarshal(jsonBytes, &zapCfg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 创建自定义日志对象
	zapLogger, err := zapCfg.Build()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer zapLogger.Sync()

	logger = zapLogger
}

// Logger 获取 logger 对象
func Logger() *zap.Logger {
	return logger
}

// 不使用接口进行初始化，可以在各个对象中定义 Init 函数在 init 包中调用。
// Init 初始化日志对象
// func Init() {
// 	logJSON := []byte(`{
// 	"level":"debug",
// 	"encoding": "json",
// 	"outputPaths": ["stdout", "/tmp/logs"],
// 	"errorOutputPaths": ["stderr"],
// 	"initialFields": {"foo": "bar"},
// 	"encoderConfig": {
// 	  "messageKey": "message",
// 	  "levelKey": "level",
// 	  "levelEncoder": "lowercase"
// 	}
// 	}`)

// 	var cfg zap.Config
// 	if err := json.Unmarshal(logJSON, &cfg); err != nil {
// 		panic(err)
// 	}
// 	logger, err := cfg.Build()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer logger.Sync()

// }
