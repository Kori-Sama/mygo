package model

import (
	"fmt"
	"mygo/config"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func InitEngine() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DbName,
		config.Database.Charset)

	var err error
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	log.Infof("Succeed to connect to database %s", config.Database.DbName)
}

func SyncTables() {
	err := engine.Sync2(new(User))
	if err != nil {
		log.Warnf("Failed to sync database: %s", err)
	}
	log.Infof("Succeed to sync the tables of database %s", config.Database.DbName)
}
