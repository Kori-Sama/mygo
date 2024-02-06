package model

import (
	"fmt"
	"log"
	"mygo/config"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func InitEngine() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DbName)

	var err error
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
}

func SyncTables() {
	err := engine.Sync2(new(User))
	if err != nil {
		log.Fatalf("Failed to sync database: %s", err)
	}
}
