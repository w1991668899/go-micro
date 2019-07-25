package rpc_client

import (
	"go-micro/golib/lib/lib_config"
	pbuser "go-micro/golib/protoc/server_one"
	"go-micro/golib/rpcservice"
)

var(
	ServerOneClient *pbuser.UserService
)

func InitClient(config lib_config.ConfMicroRpcService)  {
	if config.ServiceName == "" {
		config.ServiceName = rpcservice.ApiGatewayService
	}
	service := rpcservice.CreateService(config)

	//ServerOneClient = pblogin.NewLoginService(rpcservice.LoginService, service.Client())

	ServerOneClient = pbuser.GetUserByUserNameReq{}
}
