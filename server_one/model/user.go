package model

import (
	"go-micro/golib/lib/lib_gorm"
)

type User struct {
	lib_gorm.BaseModel
	Name string
	Age int
	lib_gorm.TimeModel
}
