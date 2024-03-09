package model

import (
	"mygo/internal/pkg/common"
	"time"
)

type History struct {
	HistoryID     int           `xorm:"pk autoincr 'history_id'"`
	UserID        int           `xorm:"'user_id'"`
	TransactionID int           `xorm:"'transaction_id'"`
	Action        common.Action `xorm:"type:action 'action'"`
	Timestamp     time.Time     `xorm:"'timestamp'"`
	NewValue      string        `xorm:"'new_value'"`
}

func (h *History) ToResponse() *common.HistoryResponse {
	return &common.HistoryResponse{
		ID:            h.HistoryID,
		UserID:        h.UserID,
		TransactionID: h.TransactionID,
		Action:        h.Action,
		Timestamp:     h.Timestamp.Unix(),
		NewValue:      h.NewValue,
	}
}

func CreateHistory(userID int, transactionID int, action common.Action, newValue string) error {
	history := History{
		UserID:        userID,
		TransactionID: transactionID,
		Action:        action,
		Timestamp:     time.Now(),
		NewValue:      newValue,
	}
	_, err := engine.Insert(&history)
	return err
}

func GetHistoryById(historyID int) (*History, error) {
	history := new(History)
	_, err := engine.ID(historyID).Get(history)
	return history, err
}

func GetHistoryByUserId(userID int) ([]*History, error) {
	histories := make([]*History, 0)
	err := engine.Where("user_id = ?", userID).Find(&histories)
	return histories, err
}

func GetHistoryByTransactionId(transactionID int) ([]*History, error) {
	histories := make([]*History, 0)
	err := engine.Where("transaction_id = ?", transactionID).Find(&histories)
	return histories, err
}

func GetHistoryByAction(action common.Action) ([]*History, error) {
	histories := make([]*History, 0)
	err := engine.Where("action = ?", action).Find(&histories)
	return histories, err
}

func GetAllHistory() ([]*History, error) {
	histories := make([]*History, 0)
	err := engine.Find(&histories)
	return histories, err
}
