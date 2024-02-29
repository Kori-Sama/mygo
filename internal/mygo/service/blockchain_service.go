package service

import (
	"math/big"
	"mygo/internal/blockchain"
	"mygo/internal/mygo/model"
	"mygo/internal/pkg/common"
)

func CreateWallet(id int, passphrase string) error {
	if passphrase == "" {
		return common.ErrorEmpty
	}
	user, err := model.GetUserById(id)
	if err != nil {
		return err
	}

	wallet, err := blockchain.NewAccount(passphrase)
	if err != nil {
		return err
	}
	if err = user.UpdateWallet(wallet); err != nil {
		return err
	}
	if err = user.UpdatePassphrase(passphrase); err != nil {
		return err
	}

	return nil
}

func GetBalance(id int) (string, error) {
	user, err := model.GetUserById(id)
	if err != nil {
		return "", err
	}
	if user.Wallet == "" {
		return "", common.ErrorNoWallet
	}
	balance, err := blockchain.BalanceOf(user.Wallet)
	if err != nil {
		return "", err
	}

	return balance.String(), nil
}

func Transfer(username, passphrase, toName string, amount string) error {
	if username == "" || passphrase == "" || toName == "" || amount == "" {
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

	balance, err := blockchain.BalanceOf(user.Wallet)
	if err != nil {
		return err
	}
	amountBigInt, isValid := new(big.Int).SetString(amount, 10)
	if !isValid {
		return common.ErrorInvalidAmount
	}
	if amountBigInt.Cmp(big.NewInt(0)) == -1 || amountBigInt.Cmp(big.NewInt(0)) == 0 {
		return common.ErrorInvalidAmount
	}

	if amountBigInt.Cmp(balance) == 1 {
		return common.ErrorBalanceNotEnough
	}

	err = blockchain.Transfer(user.Wallet, passphrase, to.Wallet, amountBigInt)
	if err != nil {
		return err
	}

	return nil
}
