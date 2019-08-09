package main

import (
	"flag"
	"go-micro/server-one/common"
	"go-micro/server-one/model"
	"go-micro/server-one/mongodb"
	"go-micro/server-one/rpc_server"
)

func main()  {
	configPath := flag.String("c", "./config/dev.yaml", "full path config file")
	flag.Parse()

	common.InitConfig(*configPath)
	model.InitModel(common.GloConf.DB)
	mongodb.InitMongo(common.GloConf.MongoDb)

	rpc_server.Start(common.GloConf.Micro)
}
