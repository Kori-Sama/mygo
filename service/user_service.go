package service

import (
	"mygo/model"
	"mygo/pkg/common"
)

func Login(username string, password string) error {
	if username == "" || password == "" {
		return common.ErrorEmpty
	}

	user, err := model.GetUserByName(username)
	if err != nil {
		return err
	}

	if user.Password != password {
		return common.ErrorWrongPassword
	}
	return nil
}
