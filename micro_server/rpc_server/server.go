package rpc_server

import (
	"go-micro/golib/lib_config"
	"go-micro/golib/lib_rpc_service"
	"log"
)

func Start(microConf lib_config.ConfMicroRpcService)  {
	if microConf.ServiceName == "" {
		log.Fatalln("service name can not be null")
	}
	service := lib_rpc_service.CreateService(microConf)
	if err := service.Run(); err != nil{
		log.Fatalln("service.Run() fail: ", err.Error())
	}
}