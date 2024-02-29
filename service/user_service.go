package service

import (
	"mygo/model"
	"mygo/pkg/common"
	"mygo/pkg/utils"
)

func Login(username string, password string) (int, string, error) {
	if username == "" || password == "" {
		return 0, "", common.ErrorEmpty
	}

	user, err := model.GetUserByName(username)
	if err != nil {
		return 0, "", err
	}

	if !utils.CmpPwd(user.Password, password) {
		return 0, "", common.ErrorWrongPassword
	}

	return user.Id, user.Role, nil
}

func Register(username string, password string, role string) (int, error) {
	if username == "" || password == "" {
		return 0, common.ErrorEmpty
	}

	if _, err := model.GetUserByName(username); err == nil {
		return 0, common.ErrorRepeatUsername
	}
	encryptedPwd := utils.Encrypt(password)
	return model.CreateUser(username, encryptedPwd, role)
}
