package modle

import (
	"github.com/jinzhu/gorm"
)

// Product table
type Product struct {
	gorm.Model // gorm 的默认会创建的结构
	Code       string
	Price      uint
}
