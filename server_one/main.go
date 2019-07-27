package main

import (
	"flag"
	"go-micro/server_one/common"
	"go-micro/server_one/model"
	"go-micro/server_one/rpc_server"
)

func main()  {
	configPath := flag.String("c", "./config/dev.yaml", "full path config file")
	flag.Parse()

	common.InitConfig(*configPath)
	model.InitModel(common.GloConf.DB)
	rpc_server.Start(common.GloConf.Micro)
}
