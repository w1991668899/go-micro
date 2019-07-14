package common

import "go-micro/golib/lib_config"

type Config struct {
	Log    lib_config.ConfLog             `yaml:"log"`
	DB     lib_config.ConfMysql           `yaml:"db"`
	//Redis  lib_config.ConfRedis           `yaml:"redis"`
	//Rabbit lib_config.ConfRabbitMQ        `yaml:"rabbit"`
	Micro  lib_config.ConfMicroRpcService `yaml:"micro"`
	//OpenTracing lib_config.Conf    `yaml:"opentracing"`
}
