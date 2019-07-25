package api

import (
	"context"
	pbuser "go-micro/golib/protoc/server_one"
	"go-micro/server_one/dao"
	"go-micro/server_one/model"
)

func GetUser(ctx context.Context, req *pbuser.UserName, resp *pbuser.User) error {
	userM := &model.User{}
	userM.Name = req.Name
	userMSli, err := dao.GetUser(*userM)
	if err != nil{
		return err
	}

	resp.Name = userMSli[0].Name
	resp.Id = int32(userMSli[0].Id)
	resp.Age = int32(userMSli[0].Age)
	resp.CreateAt = userMSli[0].CreatedAt.String()
	resp.UpdateAt = userMSli[0].UpdatedAt.String()

	return nil
}
