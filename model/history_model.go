package model

import (
	"mygo/pkg/common"
	"time"
)

type History struct {
	HistoryID     int           `xorm:"pk autoincr 'history_id'"`
	UserID        int           `xorm:"FOREIGN KEY (user_id) REFERENCES user(id)"`
	TransactionID int           `xorm:"FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id)"`
	Action        common.Action `xorm:"enum('Create', 'Edit', 'Delete', 'Respond') notnull 'action'"`
	Timestamp     time.Time     `xorm:"'timestamp'"`
}
