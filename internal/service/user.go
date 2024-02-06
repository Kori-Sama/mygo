package service

import (
	"mygo/internal/model"
)

func Login(username string, password string) error {
	if username == "" || password == "" {
		return ErrorEmpty
	}

	user, err := model.GetUserByName(username)
	if err != nil {
		return ErrorUnknownUsername
	}

	if user.Password != password {
		return ErrorWrongPassword
	}
	return nil
}
