package model

import (
	"time"
)

type Transaction struct {
	TransactionID int       `xorm:"pk autoincr 'transaction_id'"`
	UserID        int       `xorm:"'user_id'"`
	Title         string    `xorm:"varchar(100) notnull 'title'"`
	Description   string    `xorm:"text notnull 'description'"`
	Status        string    `xorm:"'status'"`
	CreatedAt     time.Time `xorm:"created 'created_at'"`
	UpdatedAt     time.Time `xorm:"updated 'updated_at'"`
}

const (
	Draft     = "Draft"
	Published = "Published"
	Closed    = "Closed"
)
