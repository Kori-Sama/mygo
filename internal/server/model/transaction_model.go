package model

import (
	"mygo/config"
	"mygo/internal/pkg/common"
	"time"
)

type Transaction struct {
	TransactionID int           `xorm:"pk autoincr 'transaction_id'" json:"id"`
	UserID        int           `xorm:"'user_id'" json:"user_id"`
	Title         string        `xorm:"varchar(100) notnull 'title'" json:"title"`
	Description   string        `xorm:"text notnull 'description'" json:"description"`
	Value         int           `xorm:"notnull 'value'" json:"value"`
	Status        common.Status `xorm:"type:status 'status'" json:"status"`
	CreatedAt     time.Time     `xorm:"created 'created_at'" json:"created_at"`
	UpdatedAt     time.Time     `xorm:"updated 'updated_at'" json:"updated_at"`
}

func (t *Transaction) ToResponse() *common.TransactionResponse {
	return &common.TransactionResponse{
		ID:          t.TransactionID,
		UserID:      t.UserID,
		Title:       t.Title,
		Description: t.Description,
		Value:       t.Value,
		Status:      t.Status,
		CreatedAt:   t.CreatedAt.Unix(),
		UpdatedAt:   t.UpdatedAt.Unix(),
	}
}

func CreateTransaction(userID int, t common.NewTransactionRequest, status common.Status) (Transaction, error) {
	transaction := Transaction{
		UserID:      userID,
		Title:       t.Title,
		Description: t.Description,
		Value:       t.Value,
		Status:      status,
	}
	_, err := engine.Insert(&transaction)
	if err != nil {
		return transaction, common.ErrorOperateDatabase
	}
	return transaction, nil
}

func GetTransactionById(id int) (*Transaction, error) {
	transaction := &Transaction{}
	isFind, err := engine.ID(id).Get(transaction)
	if err != nil {
		return nil, common.ErrorOperateDatabase
	}
	if !isFind {
		return nil, common.ErrorUnknownTransaction
	}
	return transaction, nil
}

func (t *Transaction) Update() error {
	_, err := engine.ID(t.TransactionID).Update(t)
	if err != nil {
		return common.ErrorOperateDatabase
	}
	return nil
}

func UpdateTransaction(id int, t common.TransactionRequest, status common.Status) (Transaction, error) {
	transaction := &Transaction{
		TransactionID: id,
		Title:         t.Title,
		Description:   t.Description,
		Value:         t.Value,
		Status:        status,
	}
	return *transaction, transaction.Update()
}

func GetTransactionsByUserID(userID int) ([]*Transaction, error) {
	transactions := make([]*Transaction, 0)
	err := engine.Where("user_id = ?", userID).Find(&transactions)
	if err != nil {
		return nil, common.ErrorOperateDatabase
	}
	return transactions, nil
}

func GetTransactionsByStatus(status common.Status) ([]*Transaction, error) {
	transactions := make([]*Transaction, 0)
	err := engine.Where("status = ?", status).Find(&transactions)
	if err != nil {
		return nil, common.ErrorOperateDatabase
	}
	return transactions, nil
}

func GetTransactionsByUserIDAndStatus(userID int, status common.Status) ([]*Transaction, error) {
	transactions := make([]*Transaction, 0)
	err := engine.Where("user_id = ? and status = ?", userID, status).Find(&transactions)
	if err != nil {
		return nil, common.ErrorOperateDatabase
	}
	return transactions, nil
}

func GetAllTransactions() ([]*Transaction, error) {
	transactions := make([]*Transaction, 0)
	err := engine.Find(&transactions)
	if err != nil {
		return nil, common.ErrorOperateDatabase
	}
	return transactions, nil
}

func GetPassedTransactions() ([]*Transaction, error) {
	transactions := make([]*Transaction, 0)
	err := engine.Where("status = ?", common.StatusPassed).Find(&transactions)
	if err != nil {
		return nil, common.ErrorOperateDatabase
	}
	return transactions, nil
}

func DeleteTransaction(id int) error {
	_, err := engine.ID(id).Delete(new(Transaction))
	if err != nil {
		return common.ErrorOperateDatabase
	}
	return nil
}

func GetLimitedTransactions(offset int) ([]*Transaction, error) {
	limit := config.Database.Limit
	if limit == 0 {
		limit = 20
	}
	offset = offset * limit
	transactions := make([]*Transaction, 0)
	err := engine.Limit(limit, offset).Find(&transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func GetLimitedPassedTransactions(offset int) ([]*Transaction, error) {
	limit := config.Database.Limit
	if limit == 0 {
		limit = 20
	}
	offset = offset * limit
	transactions := make([]*Transaction, 0)
	err := engine.Limit(limit, offset).Where("status = ?", common.StatusPassed).Find(&transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func GetAllTransactionCount() (int64, error) {
	count, err := engine.Count(new(Transaction))
	if err != nil {
		return -1, err
	}
	return count, nil
}

func GetPassedTransactionCount() (int64, error) {
	count, err := engine.Where("status = ?", common.StatusPassed).Count(new(Transaction))
	if err != nil {
		return -1, err
	}
	return count, nil
}
