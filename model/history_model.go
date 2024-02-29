package model

import (
	"time"
)

type History struct {
	HistoryID     int       `xorm:"pk autoincr 'history_id'"`
	UserID        int       `xorm:"'user_id'"`
	TransactionID int       `xorm:"'transaction_id'"`
	Action        string    `xorm:"'action'"`
	Timestamp     time.Time `xorm:"'timestamp'"`
}

const (
	Create  = "Create"
	Edit    = "Edit"
	Delete  = "Delete"
	Respond = "Respond"
)
