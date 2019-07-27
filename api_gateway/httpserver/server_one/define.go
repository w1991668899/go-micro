package server_one

import (
	"github.com/labstack/echo"
	"go-micro/golib/toolkit/tool_http"
)

func MountUserApi(g *echo.Group) {
	userGroup := g.Group("/server_one")
	userGroup.GET("/user/:name", tool_http.EchoResponseWrapper(HttpGetUser))
}
