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
	Status        common.Status `xorm:"enum('Draft', 'Published','Censoring','Passed') default 'Draft' 'status'"`
	CreatedAt     time.Time     `xorm:"created 'created_at'"`
	UpdatedAt     time.Time     `xorm:"updated 'updated_at'"`
}
