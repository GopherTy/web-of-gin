package logger

import (
	"encoding/json"

	"go.uber.org/zap"
)

// Init 初始化日志对象
func Init() {
	logJSON := []byte(`{
	"level":"debug", 
	"encoding": "json",
	"outputPaths": ["stdout", "/tmp/logs"],
	"errorOutputPaths": ["stderr"],
	"initialFields": {"foo": "bar"},
	"encoderConfig": {
	  "messageKey": "message",
	  "levelKey": "level",
	  "levelEncoder": "lowercase"
	}
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(logJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

}
