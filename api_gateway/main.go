package main

import (
	"flag"
	"go-micro/api_gateway/common"
	"go-micro/api_gateway/httpserver"
	"go-micro/api_gateway/rpc_client"
	"go-micro/golib/lib/lib_middleware/opentracing"
)

func main(){
	configPath := flag.String("c", "./config/dev.yaml", "full path config file")
	flag.Parse()

	common.InitConfig(*configPath)

	_, closer := opentracing.NewTracerByConfig(common.GloConf.OpenTracing)
	defer closer.Close()

	rpc_client.InitClient(common.GloConf.Micro)

	httpserver.StartHttpServer(common.GloConf.Http)
}
