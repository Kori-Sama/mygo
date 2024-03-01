package service

import (
	"mygo/internal/pkg/common"
	"mygo/internal/server/model"
)

func GetTransaction(userID, transactionID int) (*common.TransactionResponse, error) {
	transaction, err := model.GetTransactionById(transactionID)
	if err != nil {
		return nil, err
	}
	if transaction.UserID != userID {
		return nil, common.ErrorMatchTransaction
	}
	return &common.TransactionResponse{
		ID:          transaction.TransactionID,
		Title:       transaction.Title,
		Description: transaction.Description,
		Value:       transaction.Value,
		Status:      transaction.Status,
		CreatedAt:   transaction.CreatedAt.Second(),
		UpdatedAt:   transaction.UpdatedAt.Second(),
	}, nil
}

func GetTransactions(userID int) ([]common.TransactionResponse, error) {
	transactions, err := model.GetTransactionsByUserID(userID)
	if err != nil {
		return nil, err
	}
	var res []common.TransactionResponse
	for _, t := range transactions {
		res = append(res, common.TransactionResponse{
			ID:          t.TransactionID,
			Title:       t.Title,
			Description: t.Description,
			Value:       t.Value,
			Status:      t.Status,
			CreatedAt:   t.CreatedAt.Second(),
			UpdatedAt:   t.UpdatedAt.Second(),
		})
	}
	return res, nil
}

func NewTransaction(userID int, transaction common.NewTransactionRequest) (int, error) {
	return model.CreateTransaction(userID, transaction, common.StatusDraft)
}

func SaveTransaction(userID int, transaction common.TransactionRequest) error {
	return updateTransaction(userID, transaction, common.StatusDraft)
}

// Published transaction will be censored by admin
func PublishTransaction(userID int, transaction common.TransactionRequest) error {
	return updateTransaction(userID, transaction, common.StatusCensoring)
}

func updateTransaction(userID int, transaction common.TransactionRequest, status common.Status) error {
	_, err := model.GetUserById(userID)
	if err != nil {
		return err
	}
	t, err := model.GetTransactionById(transaction.ID)
	if err != nil {
		return err
	}
	if t.UserID != userID {
		return common.ErrorMatchTransaction
	}
	err = model.UpdateTransaction(t.TransactionID, transaction, status)
	if err != nil {
		return err
	}
	return nil
}
