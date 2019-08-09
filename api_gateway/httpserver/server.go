package httpserver

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-micro/api_gateway/common"
	"go-micro/api_gateway/httpserver/server_one"
	mid "go-micro/api_gateway/middle"
	"go-micro/golib/lib/lib_config"
	"go-micro/golib/lib/lib_middleware/metrics/prometheus"
	"go-micro/golib/lib/lib_middleware/opentracing"
	"go-micro/golib/toolkit/tool_http"
)

func StartHttpServer(httpConfig lib_config.ConfHttp)  {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(common.GloConf.Cors))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: common.LibLog.Logger.Out}))
	e.Use(opentracing.OpenTracing())
	e.Use(prometheus.APIMetricsByConfig(common.GloConf.Metrics))
	//e.Use(mid.StatusAuth())
	e.Use(mid.QueryParamsCheck())
	e.HTTPErrorHandler = tool_http.ServerErrorHandler

	apiGroup := e.Group("/api/v1")
	server_one.MountUserApi(apiGroup)

	// 启动http 服务
	e.Logger.Fatal(e.Start(httpConfig.Host))
}
