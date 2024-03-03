package service

import (
	"mygo/internal/pkg/common"
	"mygo/internal/server/model"
	"sort"

	"github.com/go-ego/gse"
)

// Global variable Declaration

var Seg gse.Segmenter

// End of Global variable Declaration

func SearchTransactions(search string) ([]common.TransactionResponse, error) {
	transactions, err := model.GetPassedTransactions()
	if err != nil {
		return nil, err
	}
	var res []common.TransactionResponse

	segment := Seg.CutSearch(search, true)
	searchMap := make(map[string]int)

	for _, s := range segment {
		value, _, _ := Seg.Value(s) // should have some err handling
		searchMap[s] = value
		// if err != nil {
		// 	return nil, err
		// }
	}

	searchResult := [][2]any{} // [transaction, score]
	for _, t := range transactions {
		score := 0
		titleSeg := Seg.CutSearch(t.Title, true)
		for _, s := range titleSeg {
			if _, ok := searchMap[s]; ok {
				score += 3
			}
		}

		descriptionSeg := Seg.CutSearch(t.Description, true)
		for _, s := range descriptionSeg {
			if _, ok := searchMap[s]; ok {
				score += 1
			}
		}

		if score != 0 {
			searchResult = append(searchResult, [2]any{t, score})
		}
	}
	sort.Slice(searchResult, func(i, j int) bool {
		return searchResult[i][1].(int) > searchResult[j][1].(int)
	})

	for _, r := range searchResult {
		res = append(res, *r[0].(*model.Transaction).ToResponse())
	}

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

func CensorTransaction(isPassed bool, transactionID int) error {
	t, err := model.GetTransactionById(transactionID)
	if err != nil {
		return err
	}

	if isPassed {
		t.Status = common.StatusPassed
	} else {
		t.Status = common.StatusRejected
	}
	err = t.Update()

	if err != nil {
		return err
	}
	return nil
}

func GetLimitedTransactions(role common.Role, limit int) ([]common.TransactionResponse, error) {
	var transactions []*model.Transaction
	var err error
	if role == common.RoleAdmin {
		transactions, err = model.GetLimitedTransactions(limit)
	} else {
		transactions, err = model.GetLimitedPassedTransactions(limit)
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
