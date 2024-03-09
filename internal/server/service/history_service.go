package service

import (
	"mygo/internal/pkg/common"
	"mygo/internal/server/model"
)

func CreateHistory(userID int, transactionID int, action common.Action, newValue string) error {
	return model.CreateHistory(userID, transactionID, action, newValue)
}

func GetHistoryById(historyID int) (*common.HistoryResponse, error) {
	history, err := model.GetHistoryById(historyID)
	if err != nil {
		return nil, err
	}
	return history.ToResponse(), nil
}

func GetHistoryByUserId(userID int) ([]*common.HistoryResponse, error) {
	histories, err := model.GetHistoryByUserId(userID)
	if err != nil {
		return nil, err
	}
	var historyResponses []*common.HistoryResponse
	for _, history := range histories {
		historyResponses = append(historyResponses, history.ToResponse())
	}
	return historyResponses, nil
}

func GetHistoryByTransactionId(transactionID int) ([]*common.HistoryResponse, error) {
	histories, err := model.GetHistoryByTransactionId(transactionID)
	if err != nil {
		return nil, err
	}
	var historyResponses []*common.HistoryResponse
	for _, history := range histories {
		historyResponses = append(historyResponses, history.ToResponse())
	}
	return historyResponses, nil
}

func GetHistoryByAction(action common.Action) ([]*common.HistoryResponse, error) {
	histories, err := model.GetHistoryByAction(action)
	if err != nil {
		return nil, err
	}
	var historyResponses []*common.HistoryResponse
	for _, history := range histories {
		historyResponses = append(historyResponses, history.ToResponse())
	}
	return historyResponses, nil
}

func GetAllHistory() ([]*common.HistoryResponse, error) {
	histories, err := model.GetAllHistory()
	if err != nil {
		return nil, err
	}
	var historyResponses []*common.HistoryResponse
	for _, history := range histories {
		historyResponses = append(historyResponses, history.ToResponse())
	}
	return historyResponses, nil
}
