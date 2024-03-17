package service

import (
	"mygo/internal/pkg/common"
	"mygo/internal/server/model"
	"strconv"
)

// func Transfer(senderID int, receiverID int, amount int) (*model.Transfer, error) {
// 	if senderID == receiverID {
// 		return nil, common.ErrorSameID
// 	}
// 	if amount <= 0 {
// 		return nil, common.ErrorNegativeAmount
// 	}
// 	sender, err := model.GetUserById(senderID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	receiver, err := model.GetUserById(receiverID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	senderBalance, err := strconv.Atoi(sender.Wallet)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if senderBalance < amount {
// 		return nil, common.ErrorInsufficientBalance
// 	}

// 	receiverBalance, err := strconv.Atoi(receiver.Wallet)
// 	if err != nil {
// 		return nil, err
// 	}

// 	senderBalance -= amount
// 	receiverBalance += amount

// 	sender.Wallet = strconv.Itoa(senderBalance)
// 	receiver.Wallet = strconv.Itoa(receiverBalance)

// 	newTranfer, err := model.CreateTransfer(sender, receiver, amount)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return newTranfer, nil
// }

func Transfer(senderName, passphrase, receiverName string, amount string) error {
	sender, err := model.GetUserByName(senderName)
	if err != nil {
		return err
	}
	if sender.Passphrase != passphrase {
		return common.ErrorWrongPassphrase
	}

	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		return common.ErrorInvalidAmount
	}

	if sender.Wallet < amountInt {
		return common.ErrorInsufficientBalance
	}

	receiver, err := model.GetUserByName(receiverName)
	if err != nil {
		return err
	}

	sender.Wallet -= amountInt
	receiver.Wallet += amountInt

	_, err = model.CreateTransfer(sender, receiver, amountInt)
	if err != nil {
		return err
	}

	return nil
}

func CreateWallet(userID int, passphrase string) error {
	user, err := model.GetUserById(userID)
	if err != nil {
		return err
	}
	if user.Wallet != 0 {
		return common.ErrorRepeatUsername
	}
	err = user.UpdateWallet("100") // default wallet
	if err != nil {
		return err
	}
	err = user.UpdatePassphrase(passphrase)
	if err != nil {
		return err
	}
	return nil
}

func GetBalance(userID int) (int, error) {
	user, err := model.GetUserById(userID)
	if err != nil {
		return 0, err
	}
	return user.Wallet, nil
}

func GetTransfersByID(transferID int) (*common.TransferHistoryResponse, error) {
	transfer, err := model.GetTransferById(transferID)
	if err != nil {
		return nil, err
	}
	return transfer.ToResponse(), nil
}

func GetTransfersBySenderID(senderID int) ([]*common.TransferHistoryResponse, error) {
	transfers, err := model.GetTransferBySenderId(senderID)
	if err != nil {
		return nil, err
	}
	var transferResponses []*common.TransferHistoryResponse
	for _, transfer := range transfers {
		transferResponses = append(transferResponses, transfer.ToResponse())
	}
	return transferResponses, nil
}

func GetTransfersByReceiverID(receiverID int) ([]*common.TransferHistoryResponse, error) {
	transfers, err := model.GetTransferByReceiverId(receiverID)
	if err != nil {
		return nil, err
	}
	var transferResponses []*common.TransferHistoryResponse
	for _, transfer := range transfers {
		transferResponses = append(transferResponses, transfer.ToResponse())
	}
	return transferResponses, nil
}
