package service

import (
	"mygo/blockchain"
	"mygo/model"
	"mygo/pkg/common"
)

func CreateWallet(username, passphrase string) error {
	if username == "" || passphrase == "" {
		return common.ErrorEmpty
	}
	user, err := model.GetUserByName(username)
	if err != nil {
		return err
	}

	wallet, err := blockchain.NewAccount(passphrase)
	if err != nil {
		return common.ErrorCreateWallet
	}
	if err = user.UpdateWallet(wallet); err != nil {
		return common.ErrorOperateDatabase
	}

	return nil
}
