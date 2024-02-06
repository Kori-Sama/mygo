package service

import (
	"mygo/internal/blockchain"
	"mygo/internal/model"
	utils "mygo/internal/pkg"
)

func CreateWallet(username, passphrase string) error {
	if username == "" || passphrase == "" {
		return utils.ErrorEmpty
	}
	user, err := model.GetUserByName(username)
	if err != nil {
		return err
	}

	wallet, err := blockchain.NewAccount(passphrase)
	if err != nil {
		return utils.ErrorCreateWallet
	}
	if err = user.UpdateWallet(wallet); err != nil {
		return utils.ErrorOperateDatabase
	}

	return nil
}
