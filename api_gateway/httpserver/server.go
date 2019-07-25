package httpserver

import (
	"github.com/gin-gonic/gin"
	"go-micro/api_gateway/httpserver/api/server_one"
)

func StartHttpServer()  {
	r := gin.New()

	r.Use(gin.Recovery())
	//r.Use()

	apiGroup := r.Group("/api/v1/")

	apiGroup.GET("name", )

	server_one.MountUserApi(apiGroup)
}
