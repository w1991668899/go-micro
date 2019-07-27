package rpc_server

import (
	"go-micro/golib/lib/lib_config"
	pbserverone "go-micro/golib/protoc/server_one"
	"go-micro/golib/rpcservice"
	"go-micro/server_one/rpc_server/api"
	"log"
)

func Start(microConf lib_config.ConfMicroRpcService)  {
	if microConf.ServiceName == "" {
		microConf.ServiceName = rpcservice.ServerOneService
	}

	service := rpcservice.CreateService(microConf)
	pbserverone.RegisterServerOneServiceHandler(service.Server(), new(api.User))

	if err := service.Run(); err != nil{
		log.Fatalln("service.Run() fail: ", err.Error())
	}
}