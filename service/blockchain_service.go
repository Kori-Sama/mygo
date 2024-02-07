package service

import (
	"math"
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

func Transfer(username, passphrase, toName string, amount float64) error {
	if username == "" || passphrase == "" || toName == "" || amount <= 0 {
		return common.ErrorEmpty
	}
	user, err := model.GetUserByName(username)
	if err != nil {
		return err
	}
	to, err := model.GetUserByName(toName)
	if err != nil {
		return err
	}
	if user.Wallet == "" || to.Wallet == "" {
		return common.ErrorNoWallet
	}

	decimal, err := blockchain.Decimal()
	if err != nil {
		return err
	}
	balance, err := blockchain.BalanceOf(user.Wallet)
	if err != nil {
		return err
	}

	if amount > calcToken(balance, decimal) {
		return common.ErrorBalanceNotEnough
	}

	amountBigInt := new(big.Int).SetInt64(int64(math.Pow(amount, float64(decimal))))
	err = blockchain.Transfer(user.Wallet, passphrase, to.Wallet, amountBigInt)
	if err != nil {
		return err
	}

	return nil
}

func calcToken(balance *big.Int, decimal uint8) float64 {
	b := new(big.Float).SetInt(balance)
	d := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil))
	f, _ := new(big.Float).Quo(b, d).Float64()
	return f
}
