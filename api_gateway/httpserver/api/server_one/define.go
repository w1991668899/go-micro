package server_one

import "github.com/gin-gonic/gin"

func MountUserApi(g *gin.RouterGroup) {
	userGroup := g.Group("server_one/")

	userGroup.GET("name", )

}
