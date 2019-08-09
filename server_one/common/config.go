package common

import (
	"go-micro/golib/lib/lib_config"
	"go-micro/golib/lib/lib_log"
)

var (
    GloConf *Config
    LibLog *lib_log.LibLog
)

func InitConfig(configFilePath... string)  {
	GloConf = &Config{}
	lib_config.LoadConfig(GloConf, configFilePath...)
	LibLog = lib_log.InitLog(GloConf.Log)
}
