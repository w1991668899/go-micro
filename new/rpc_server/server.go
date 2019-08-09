package rpc_server

import (
	"github.com/sirupsen/logrus"
	"go-micro/golib/lib/lib_config"
	pbserverone "go-micro/golib/protoc/server_one"
	"go-micro/golib/rpcservice"
	"go-micro/server-one/common"
	"go-micro/server-one/rpc_server/api"
)

func Start(microConf lib_config.ConfMicroRpcService)  {
	if microConf.ServiceName == "" {
		microConf.ServiceName = rpcservice.ServerOneService
	}

	service := rpcservice.CreateService(microConf)
	err := pbserverone.RegisterServerOneServiceHandler(service.Server(), new(api.User))
	if err != nil{
		common.LibLog.LogPanic(logrus.Fields{
			"err": err,
		}, "service registry fail")
	}

	if err = service.Run(); err != nil{
		common.LibLog.LogPanic(logrus.Fields{
			"err": err,
		}, "service.Run fail")
	}

	//s := web.NewService(web.Name(rpcservice.ServerOneService), web.Registry(etcdv3.NewRegistry(registry.Addrs(common.GloConf.Micro.EtcdAddrSli[0]))))
	//if err := s.Init(); err != nil {
	//	log.Fatal(err)
	//}
	//
	//if err := service.Run(); err != nil {
	//	log.Fatal(err)
	//}

}