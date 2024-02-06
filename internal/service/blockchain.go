package service

import (
	"errors"
	"mygo/internal/blockchain"
	"mygo/internal/model"
)

func CreateWallet(username, passphrase string) error {
	if passphrase == "" {
		return errors.New("passphrase cannot be empty")
	}
	user, err := model.GetUserByName(username)
	if err != nil {
		return ErrorUnknownUsername
	}

	wallet, err := blockchain.NewAccount(passphrase)
	if err != nil {
		return ErrorCreateWallet
	}
	if err = user.UpdateWallet(wallet); err != nil {
		return ErrorOperateDatabase
	}

	return nil
}
