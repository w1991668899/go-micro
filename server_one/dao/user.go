package dao

import "go-micro/server_one/model"

func GetUser(user model.User) ([]*model.User, error){
	userMSli := make([]*model.User, 0)
	err := model.Db.Where(user).Find(&userMSli).Error
	return userMSli, err
}
