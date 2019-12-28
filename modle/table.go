package modle

import (
	"github.com/jinzhu/gorm"
)

// Product table
type Product struct {
	gorm.Model
	Code  string
	Price uint
}
