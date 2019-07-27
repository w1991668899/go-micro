package common

import (
	"github.com/sirupsen/logrus"
	"go-micro/golib/lib/lib_config"
	"go-micro/golib/lib/lib_log"
)

var (
    GloConf *config
    Logger *logrus.Logger
)

func InitConfig(configFilePath... string)  {
	GloConf = &config{}
	lib_config.LoadConfig(GloConf, configFilePath...)
	Logger = lib_log.InitLog(GloConf.Log)
}
