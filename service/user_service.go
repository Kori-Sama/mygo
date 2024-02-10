package service

import (
	"mygo/model"
	"mygo/pkg/common"
)

func Login(username string, password string) (int, error) {
	if username == "" || password == "" {
		return 0, common.ErrorEmpty
	}

	user, err := model.GetUserByName(username)
	if err != nil {
		return 0, err
	}

	if user.Password != password {
		return 0, common.ErrorWrongPassword
	}

	return user.Id, nil
}

func Register(username string, password string) (int, error) {
	if username == "" || password == "" {
		return 0, common.ErrorEmpty
	}

	if _, err := model.GetUserByName(username); err == nil {
		return 0, common.ErrorRepeatUsername
	}

	return model.CreateUser(username, password)
}
