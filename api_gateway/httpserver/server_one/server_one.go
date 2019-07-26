package server_one

import (
	"github.com/labstack/echo"
	"go-micro/api_gateway/rpc_api/server_one"
	"go-micro/golib/lib/lib_middleware/opentracing"
	pbserverone "go-micro/golib/protoc/server_one"
)

func HttpGetUser(ctx echo.Context) (interface{}, error) {
	req := &pbserverone.GetUserByUserNameReq{}
	req.Name = ctx.Param("name")

	context := opentracing.ContextFromEcho(ctx)
	return server_one.GetUser(context, req)
}
