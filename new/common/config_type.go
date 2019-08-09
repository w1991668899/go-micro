package common

import "go-micro/golib/lib/lib_config"

type Config struct {
	Log   lib_config.ConfLog             `yaml:"log"`
	DB    lib_config.ConfMysql           `yaml:"db"`
	Redis lib_config.ConfRedis           `yaml:"redis"`
	Micro lib_config.ConfMicroRpcService `yaml:"micro"`
	MongoDb lib_config.ConfMongoDb		 `yaml:"mongo_db"`
	//Rabbit lib_config.ConfRabbitMQ        `yaml:"rabbit"`
	//OpenTracing lib_config.Conf    `yaml:"opentracing"`
}
