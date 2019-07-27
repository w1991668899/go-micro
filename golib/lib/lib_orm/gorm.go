package lib_orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/prometheus/common/log"
	"go-micro/golib/lib/lib_config"
	"go-micro/golib/lib/lib_log"
	"go-micro/golib/lib/lib_middleware/opentracing"
	"time"
)

type BaseModel struct {
	Id int64 `gorm:"PRIMARY_KEY;Column:id;AUTO_INCREMENT"`
}

type TimeModel struct {
	UpdatedAt time.Time `gorm:"Column:updated_at"`
	CreatedAt time.Time `gorm:"Column:created_at"`
}

func CreateDb(configDb lib_config.ConfMysql) *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC",
			configDb.Username,
			configDb.Password,
			configDb.Host,
			configDb.Port,
			configDb.DBName,
	), )

	if err != nil {
		log.Fatalln("mysql connect fail: ", err)
	}

	db.LogMode(configDb.EnableLog)
	if configDb.EnableLog && configDb.LogType == "logrus" {
		db.SetLogger(&lib_log.GormLogger{})
	}

	db.SingularTable(true)

	if configDb.MaxIdle <= 0 || configDb.MaxConn <= 0{
		log.Fatalln("mysql maxidle or maxconn is fail")
	}


	db.DB().SetMaxIdleConns(configDb.MaxIdle)
	db.DB().SetMaxOpenConns(configDb.MaxConn)
	db.DB().SetConnMaxLifetime(time.Duration(configDb.MaxLifeTime) * time.Second)
	if configDb.AutoMigrate {
		db.AutoMigrate()
	}

	opentracing.SetGORMCallbacks(db)

	return db
}