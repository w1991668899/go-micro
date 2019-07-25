package model

import (
	"go-micro/golib/lib/lib_orm"
)

type User struct {
	lib_orm.BaseModel
	Name string	`json:"name"`
	Age int	`json:"age"`
	lib_orm.TimeModel
}
