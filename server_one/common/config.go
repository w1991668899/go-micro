package common

import (
	"github.com/sirupsen/logrus"
	"go-micro/golib/lib/lib_config"
	"go-micro/golib/lib/lib_log"
)

var (
    GloConf *Config
    Logger *logrus.Logger
)

func InitConfig(configFilePath... string)  {
	GloConf = &Config{}
	lib_config.LoadConfig(GloConf, configFilePath...)
	Logger = lib_log.InitLog(GloConf.Log)
}
