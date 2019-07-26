package common

import (
	"go-micro/golib/lib/lib_config"
	"go-micro/golib/lib/lib_middleware/opentracing"
)

type config struct {
	Log   lib_config.ConfLog             `yaml:"tool_log"`
	Micro lib_config.ConfMicroRpcService `yaml:"micro"`
	Http  lib_config.ConfHttp            `yaml:"http"`
	Redis  lib_config.ConfRedis           `yaml:"redis"`
	OpenTracing opentracing.ConfigJaeger   `yaml:"open_tracing"`
}
