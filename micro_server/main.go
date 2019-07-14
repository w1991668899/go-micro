package main

import (
	"flag"
	"go-micro/micro_server/common"
	"go-micro/micro_server/rpc_server"
)

func main()  {
	configPath := flag.String("c", "./config/dev.yaml", "full path config file")
	flag.Parse()
	common.InitConfig(*configPath)
	rpc_server.Start(common.GloConf.Micro)
}
