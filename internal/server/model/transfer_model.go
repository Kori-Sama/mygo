package model

import (
	"mygo/internal/pkg/common"
	"time"
)

type Transfer struct {
	TransferID int       `xorm:"pk autoincr 'transfer_id'"`
	SenderID   int       `xorm:"'sender_id'"`
	ReceiverID int       `xorm:"'receiver_id'"`
	Amount     int       `xorm:"'amount'"`
	TimeStamp  time.Time `xorm:"'timestamp'"`
}

func (t *Transfer) ToResponse() *common.TransferHistoryResponse {
	return &common.TransferHistoryResponse{
		ID:         t.TransferID,
		SenderID:   t.SenderID,
		ReceiverID: t.ReceiverID,
		Amount:     t.Amount,
		TimeStamp:  t.TimeStamp.Unix(),
	}
}

func CreateTransfer(sender *User, receiver *User, amount int) (*Transfer, error) {
	transfer := Transfer{
		SenderID:   sender.ID,
		ReceiverID: receiver.ID,
		Amount:     amount,
		TimeStamp:  time.Now(),
	}

	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return nil, err
	}

	_, err = session.ID(sender.ID).Cols("wallet").Update(sender)
	if err != nil {
		session.Rollback()
		return nil, err
	}
	_, err = session.ID(receiver.ID).Cols("wallet").Update(receiver)
	if err != nil {
		session.Rollback()
		return nil, err
	}
	_, err = session.Insert(&transfer)
	if err != nil {
		session.Rollback()
		return nil, err
	}
	err = session.Commit()
	if err != nil {
		return nil, err
	}
	return &transfer, nil
}

func GetTransferById(transferID int) (*Transfer, error) {
	transfer := new(Transfer)
	_, err := engine.ID(transferID).Get(transfer)
	return transfer, err
}

func GetTransferBySenderId(senderID int) ([]*Transfer, error) {
	transfers := make([]*Transfer, 0)
	err := engine.Where("sender_id = ?", senderID).Find(&transfers)
	return transfers, err
}

func GetTransferByReceiverId(receiverID int) ([]*Transfer, error) {
	transfers := make([]*Transfer, 0)
	err := engine.Where("receiver_id = ?", receiverID).Find(&transfers)
	return transfers, err
}

func GetAllTransfer() ([]*Transfer, error) {
	transfers := make([]*Transfer, 0)
	err := engine.Find(&transfers)
	return transfers, err
}
