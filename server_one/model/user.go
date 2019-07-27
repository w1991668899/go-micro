package model

import (
	"go-micro/golib/lib/lib_orm"
)

type User struct {
	lib_orm.BaseModel
	Name string
	Age int
	lib_orm.TimeModel
}
