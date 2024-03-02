package model

import (
	"fmt"
	"mygo/config"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func InitEngine() {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DbName,
		config.Database.SSL)

	var err error
	engine, err = xorm.NewEngine("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	log.Infof("Succeed to connect to database %s", config.Database.DbName)
}

func SyncTables() {
	err := initEnum()
	if err != nil {
		log.Warnf("Failed to init enum: %s", err)
	}

	err = engine.Sync2(new(User), new(Transaction), new(History))
	if err != nil {
		log.Warnf("Failed to sync database: %s", err)
	}
	log.Infof("Succeed to sync the tables of database %s", config.Database.DbName)
}

func initEnum() error {
	_, err := engine.Exec("create type role as enum('Old', 'Volunteer', 'Admin');")
	if err != nil {
		return err
	}
	_, err = engine.Exec("create type status as enum('Draft', 'Censoring', 'Passed','Rejected');")
	if err != nil {
		return err
	}
	_, err = engine.Exec("create type action as enum('Create', 'Update', 'Delete');")
	if err != nil {
		return err
	}
	return nil
}
