package model

import (
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"go-micro/golib/lib/lib_config"
	"go-micro/golib/lib/lib_gorm"
)

var Db *gorm.DB

func InitModel(configDb lib_config.ConfMysql)  {
	Db = lib_gorm.CreateDb(configDb)
}

func CloseDb()  {
	if Db != nil{
		err := Db.Close()
		if err != nil{
			log.Fatalln("close mysql fail: ", err)
		}
	}
}