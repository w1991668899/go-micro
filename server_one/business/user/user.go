package user

import (
	"context"
	"github.com/sirupsen/logrus"
	pbserverone "go-micro/golib/protoc/server_one"
	"go-micro/server_one/common"
	"go-micro/server_one/dao"
	"go-micro/server_one/model"
)

func GetUser(ctx context.Context, req *pbserverone.GetUserByUserNameReq, resp *pbserverone.GetUserByUserNameResp) error {
	userM := &model.User{}
	userM.Name = req.Name
	userMSli, err := dao.GetUser(*userM)
	if err != nil{
		common.LibLog.LogError(logrus.Fields{
			"err":   err,
			"func":  "GetUser",
			"param": req,
		}, "获取信息失败")
		return err
	}

	resp.Name = userMSli[0].Name
	resp.Id = int32(userMSli[0].Id)
	resp.Age = int32(userMSli[0].Age)
	resp.CreateAt = userMSli[0].CreatedAt.String()
	resp.UpdateAt = userMSli[0].UpdatedAt.String()
	return nil
}
