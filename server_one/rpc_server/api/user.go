package api

import (
	pbserverone "go-micro/golib/protoc/server_one"
	"go-micro/server_one/business/user"
	"golang.org/x/net/context"
)

type User struct {

}

func (*User) GetUser(ctx context.Context, req *pbserverone.GetUserByUserNameReq, resp *pbserverone.GetUserByUserNameResp)error  {
	return user.GetUser(ctx, req, resp)
}