package init

import (
	"web-of-gin/module"
	"web-of-gin/module/configs"
	"web-of-gin/module/db"
	"web-of-gin/module/logger"
)

// 整个项目架构为模块化
func init() {
	registers := []module.IRegister{
		configs.Register{},
		logger.Register{},
		db.Register{},
	}

	for _, r := range registers {
		r.Regist()
	}
}
