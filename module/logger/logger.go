package logger

import (
	"fmt"
	"os"
	"strings"
	"web-of-gin/module/configs"
	"web-of-gin/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Register 注册器
type Register struct {
}

var instance *zap.Logger

// Regist 实现 IRegister 接口，以注册获取初始化好的 logger 对象。
func (Register) Regist() {
	cnf := configs.Instance()

	// 是否输出日志文件
	var logPath []string
	if cnf.Logger.LogsPath != "" {
		// 创建指定路径
		err := utils.CreatePath(cnf.Logger.LogsPath)
		if err != nil {
			fmt.Println("Init logger fail: ", err)
			os.Exit(1)
		}

		_, err = os.Create(cnf.Logger.LogsPath)
		if err != nil {
			fmt.Println("Init logger fail: ", err)
			os.Exit(1)
		}

		logPath = []string{cnf.Logger.LogsPath}
	} else {
		logPath = []string{"stdout"}
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短径编码器
	}

	// 用户日志等级 debug,info,warn,error,dpanic,panic,fatal
	level := strings.TrimSpace(cnf.Logger.Level)
	level = strings.ToLower(level)
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	case "dpanic":
		zapLevel = zapcore.DPanicLevel
	case "panic":
		zapLevel = zapcore.PanicLevel
	case "fatal":
		zapLevel = zapcore.FatalLevel
	default:
		zapLevel = zapcore.DebugLevel
	}

	atomicLevel := zap.NewAtomicLevelAt(zapLevel)
	zapcnf := zap.Config{
		Level:            atomicLevel,
		Development:      cnf.Logger.Development,
		Encoding:         cnf.Logger.Encoding, // json 或 console
		OutputPaths:      logPath,
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    encoderConfig,
	}

	// 创建自定义日志对象
	zapLogger, err := zapcnf.Build()
	if err != nil {
		fmt.Println("Init logger fail: ", err)
		os.Exit(1)
	}
	defer zapLogger.Sync()

	instance = zapLogger
	instance.Info("Init logger success")
	return
}

// Instance 获取默认的 logger 对象
func Instance() *zap.Logger {
	if instance == nil {
		fmt.Println("Register logger module failed.")
		os.Exit(1)
	}
	return instance
}
