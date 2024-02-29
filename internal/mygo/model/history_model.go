package model

import (
	"mygo/internal/pkg/common"
	"time"
)

type History struct {
	HistoryID     int           `xorm:"pk autoincr 'history_id'"`
	UserID        int           `xorm:"'user_id'"`
	TransactionID int           `xorm:"'transaction_id'"`
	Action        common.Action `xorm:"enum('Create', 'Edit', 'Delete', 'Respond') notnull 'action'"`
	Timestamp     time.Time     `xorm:"'timestamp'"`
}
