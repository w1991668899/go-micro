package httpserver

import (
	"github.com/labstack/echo"
	"go-micro/api_gateway/httpserver/server_one"
	"go-micro/golib/lib/lib_config"
)

func StartHttpServer(httpConfig lib_config.ConfHttp)  {
	e := echo.New()

	//e.Use()

	apiGroup := e.Group("/api/v1")

	server_one.MountUserApi(apiGroup)

	// 启动http 服务
	e.Logger.Fatal(e.Start(httpConfig.Host))
}
