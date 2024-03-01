package model

import (
	"mygo/internal/pkg/common"
	"time"
)

type Transaction struct {
	TransactionID int           `xorm:"pk autoincr 'transaction_id'"`
	UserID        int           `xorm:"'user_id'"`
	Title         string        `xorm:"varchar(100) notnull 'title'"`
	Description   string        `xorm:"text notnull 'description'"`
	Value         int           `xorm:"notnull 'value'"`
	Status        common.Status `xorm:"enum('Draft', 'Published','Censoring','Passed','Rejected') default 'Draft' 'status'"`
	CreatedAt     time.Time     `xorm:"created 'created_at'"`
	UpdatedAt     time.Time     `xorm:"updated 'updated_at'"`
}

func CreateTransaction(userID int, title, description string, value int, status common.Status) (int, error) {
	_, err := GetUserById(userID)
	if err != nil {
		return 0, err
	}

	transaction := Transaction{
		UserID:      userID,
		Title:       title,
		Description: description,
		Value:       value,
		Status:      status,
	}
	_, err = engine.Insert(&transaction)
	if err != nil {
		return 0, common.ErrorOperateDatabase
	}
	return transaction.TransactionID, nil
}

func GetTransactionById(id int) (*Transaction, error) {
	transaction := &Transaction{}
	_, err := engine.ID(id).Get(transaction)
	if err != nil {
		return nil, common.ErrorOperateDatabase
	}
	return transaction, nil
}

func (t *Transaction) UpdateTransaction() error {
	_, err := GetUserById(t.UserID)
	if err != nil {
		return err
	}

	_, err = engine.ID(t.TransactionID).Update(t)
	if err != nil {
		return common.ErrorOperateDatabase
	}
	return nil
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

func DeleteTransaction(id int) error {
	_, err := engine.ID(id).Delete(new(Transaction))
	if err != nil {
		return common.ErrorOperateDatabase
	}
	return nil
}
