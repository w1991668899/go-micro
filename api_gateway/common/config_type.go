package common

import "go-micro/golib/lib/lib_config"

type config struct {
	Log   lib_config.ConfLog             `yaml:"log"`
	Micro lib_config.ConfMicroRpcService `yaml:"micro"`
	Http  lib_config.ConfHttp            `yaml:"http"`
	//Redis  lib_config.ConfRedis           `yaml:"redis"`
	//OpenTracing lib_config.Conf    `yaml:"opentracing"`

}
