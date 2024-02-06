package service

import (
	"mygo/internal/model"
	utils "mygo/internal/pkg"
)

func Login(username string, password string) error {
	if username == "" || password == "" {
		return utils.ErrorEmpty
	}

	user, err := model.GetUserByName(username)
	if err != nil {
		return err
	}

	if user.Password != password {
		return utils.ErrorWrongPassword
	}
	return nil
}
