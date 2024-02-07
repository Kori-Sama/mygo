package service

import (
	"math/big"
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
		return err
	}
	if err = user.UpdateWallet(wallet); err != nil {
		return common.ErrorOperateDatabase
	}

	return nil
}

func GetBalance(username string) (float64, error) {
	decimal, err := blockchain.Decimal()
	if err != nil {
		return 0, err
	}

	if username == "" {
		return 0, common.ErrorEmpty
	}
	user, err := model.GetUserByName(username)
	if err != nil {
		return 0, err
	}
	if user.Wallet == "" {
		return 0, common.ErrorNoWallet
	}
	balance, err := blockchain.BalanceOf(user.Wallet)
	if err != nil {
		return 0, err
	}

	return calcToken(balance, decimal), nil
}

func calcToken(balance *big.Int, decimal uint8) float64 {
	b := new(big.Float).SetInt(balance)
	d := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil))
	f, _ := new(big.Float).Quo(b, d).Float64()
	return f
}
