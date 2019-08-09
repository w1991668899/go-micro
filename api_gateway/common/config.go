package common

import (
	"go-micro/golib/lib/lib_config"
	"go-micro/golib/lib/lib_log"
)

var (
    GloConf *config
    LibLog *lib_log.LibLog
)

func InitConfig(configFilePath... string)  {
	GloConf = &config{}
	lib_config.LoadConfig(GloConf, configFilePath...)
	LibLog = lib_log.InitLog(GloConf.Log)
}
