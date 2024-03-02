package service

import (
	"mygo/internal/pkg/common"
	"mygo/internal/server/model"
)

func SearchTransactions(search string) ([]common.TransactionResponse, error) {
	_, err := model.GetPassedTransactions()
	if err != nil {
		return nil, err
	}
	var res []common.TransactionResponse

	// TODO 👀

	return res, nil
}

func GetTransaction(transactionID int) (*common.TransactionResponse, error) {
	t, err := model.GetTransactionById(transactionID)
	if err != nil {
		return nil, err
	}

	return t.ToResponse(), nil
}

func GetTransactionsByUserID(userID int) ([]common.TransactionResponse, error) {
	transactions, err := model.GetTransactionsByUserID(userID)
	if err != nil {
		return nil, err
	}
	var res []common.TransactionResponse
	for _, t := range transactions {
		res = append(res, *t.ToResponse())
	}
	return res, nil
}

func GetTransactions(role common.Role) ([]common.TransactionResponse, error) {
	var transactions []*model.Transaction
	var err error
	if role == common.RoleAdmin {
		transactions, err = model.GetAllTransactions()
	} else {
		transactions, err = model.GetPassedTransactions()
	}
	if err != nil {
		return nil, err
	}

	var res []common.TransactionResponse
	for _, t := range transactions {
		res = append(res, *t.ToResponse())
	}
	return res, nil
}

func GetTransactionByStatus(status common.Status) ([]common.TransactionResponse, error) {
	transactions, err := model.GetTransactionsByStatus(status)
	if err != nil {
		return nil, err
	}
	var res []common.TransactionResponse
	for _, t := range transactions {
		res = append(res, *t.ToResponse())
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

func DeleteTransaction(userID, transactionID int) error {
	t, err := model.GetTransactionById(transactionID)
	if err != nil {
		return err
	}
	if t.UserID != userID {
		return common.ErrorMatchTransaction
	}
	return model.DeleteTransaction(transactionID)
}
