package rpc_client

import (
	"go-micro/golib/lib/lib_config"
	pbserverone "go-micro/golib/protoc/server_one"
	"go-micro/golib/rpcservice"
)

var(
	ServerOneClient pbserverone.ServerOneService
)

func InitClient(configMicro lib_config.ConfMicroRpcService)  {
	if configMicro.ServiceName == "" {
		configMicro.ServiceName = rpcservice.ApiGatewayService
	}
	service := rpcservice.CreateService(configMicro)

	ServerOneClient = pbserverone.NewServerOneService(rpcservice.ServerOneService, service.Client())
}
