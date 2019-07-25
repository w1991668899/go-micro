package rpc_server

import (
	"go-micro/golib/lib/lib_config"
	"go-micro/golib/rpcservice"
	"log"
)

func Start(microConf lib_config.ConfMicroRpcService)  {
	if microConf.ServiceName == "" {
		log.Fatalln("service name can not be null")
	}
	service := rpcservice.CreateService(microConf)
	if err := service.Run(); err != nil{
		log.Fatalln("service.Run() fail: ", err.Error())
	}
}