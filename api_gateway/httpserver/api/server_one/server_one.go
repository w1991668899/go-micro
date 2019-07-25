package server_one

import (
	"github.com/gin-gonic/gin"
	pbuser "go-micro/golib/protoc/server_one"
)

func HttpGetUser(ctx gin.Context, )  {
	req := &pbuser.GetUserByUserNameReq{}
	req.Name = ctx.Param("name")

	resp := &pbuser.GetUserByUserNameResp{}

}
