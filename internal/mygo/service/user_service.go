package service

import (
	"mygo/internal/mygo/model"
	"mygo/internal/pkg/common"
	"mygo/internal/pkg/utils"
)

func Login(username string, password string) (int, common.Role, error) {
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

	return user.ID, user.Role, nil
}

func Register(username string, password string, role common.Role) (int, error) {
	if username == "" || password == "" {
		return 0, common.ErrorEmpty
	}

	if !utils.CheckRole(role) {
		return 0, common.ErrorUnknownRole
	}

	if _, err := model.GetUserByName(username); err == nil {
		return 0, common.ErrorRepeatUsername
	}
	encryptedPwd := utils.Encrypt(password)
	return model.CreateUser(username, encryptedPwd, role)
}
