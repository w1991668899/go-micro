package server_one

import (
	"context"
	"fmt"
	"go-micro/api_gateway/rpc_client"
	pbserverone "go-micro/golib/protoc/server_one"
)

func GetUser(ctx context.Context, req *pbserverone.GetUserByUserNameReq) (interface{}, error) {
	if rpc_client.ServerOneClient == nil{
		fmt.Println(req.Name, 5555555)
	}

	return rpc_client.ServerOneClient.GetUser(ctx, req)
}
