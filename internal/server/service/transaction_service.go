package service

import (
	"encoding/json"
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

func MakeDeal(userID, transactionID int) error {
	t, err := model.GetTransactionById(transactionID)
	if err != nil {
		return err
	}

	if t.Status != common.StatusPassed {
		return common.ErrorTransactionNotPassed
	}

	u, err := model.GetUserById(t.UserID)
	if err != nil {
		return err
	}
	if u.Role == common.RoleOld {
		err = Transfer(t.UserID, userID, t.Value)
	} else if u.Role == common.RoleVolunteer {
		err = Transfer(userID, t.UserID, t.Value)
	} else {
		err = common.ErrorInvalidTransaction
	}

	if err != nil {
		return err
	}

	err = model.DeleteTransaction(transactionID)
	if err != nil {
		return err
	}

	err = CreateHistory(userID, transactionID, common.ActionDelete, "")
	if err != nil {
		return err
	}
	return nil
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

func GetTransactionById(transactionID int) ([]common.TransactionResponse, error) {
	transaction, err := model.GetTransactionById(transactionID)
	if err != nil {
		return nil, err
	}
	var res []common.TransactionResponse
	res = append(res, *transaction.ToResponse())
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

func NewTransaction(userID int, transaction common.NewTransactionRequest) (model.Transaction, error) {
	newTrans, err := model.CreateTransaction(userID, transaction, common.StatusDraft)
	if err != nil {
		return model.Transaction{}, err
	}

	CreateHistory(userID, newTrans.TransactionID, common.ActionCreate, TransactionToString(newTrans))
	return newTrans, nil
}

func SaveTransaction(userID int, transaction common.TransactionRequest) (model.Transaction, error) {
	newTrans, err := updateTransaction(userID, transaction, common.StatusDraft)
	if err != nil {
		return model.Transaction{}, err
	}

	CreateHistory(userID, newTrans.TransactionID, common.ActionSave, TransactionToString(newTrans))
	return newTrans, nil
}

// Published transaction will be censored by admin
func PublishTransaction(userID int, transaction common.TransactionRequest) (model.Transaction, error) {
	newTrans, err := updateTransaction(userID, transaction, common.StatusCensoring)
	if err != nil {
		return model.Transaction{}, err
	}

	CreateHistory(userID, newTrans.TransactionID, common.ActionSave, "")
	return newTrans, nil
}

func updateTransaction(userID int, transaction common.TransactionRequest, status common.Status) (model.Transaction, error) {
	_, err := model.GetUserById(userID)
	if err != nil {
		return model.Transaction{}, err
	}
	t, err := model.GetTransactionById(transaction.ID)
	if err != nil {
		return model.Transaction{}, err
	}
	if t.UserID != userID {
		return model.Transaction{}, common.ErrorMatchTransaction
	}
	newTransaction, err := model.UpdateTransaction(t.TransactionID, transaction, status)
	if err != nil {
		return newTransaction, err
	}

	CreateHistory(userID, newTransaction.TransactionID, common.ActionEdit, TransactionToString(newTransaction))
	return newTransaction, nil
}

func DeleteTransaction(userID, transactionID int) error {
	t, err := model.GetTransactionById(transactionID)
	if err != nil {
		return err
	}
	if t.UserID != userID {
		return common.ErrorMatchTransaction
	}

	err = model.DeleteTransaction(transactionID)
	if err != nil {
		return err
	}

	CreateHistory(userID, transactionID, common.ActionDelete, "")
	return err
}

func CensorTransaction(adminID int, isPassed bool, transactionID int) error {
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

	CreateHistory(adminID, transactionID, common.ActionCensor, "")
	return nil
}

func GetLimitedTransactions(role common.Role, index int) ([]common.TransactionResponse, error) {
	var transactions []*model.Transaction
	var err error
	if role == common.RoleAdmin {
		transactions, err = model.GetLimitedTransactions(index)
	} else {
		transactions, err = model.GetLimitedPassedTransactions(index)
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

func TransactionToString(t model.Transaction) string {
	ret, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(ret)
}

func StringToTransaction(s string) *model.Transaction {
	transaction := &model.Transaction{}
	err := json.Unmarshal([]byte(s), transaction)
	if err != nil {
		return nil
	}
	return transaction
}
