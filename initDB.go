package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"
)

type Temperature struct {
	Id   int `gorm:primary_key`
	Temp float64
	Date time.Time
}

func Initdb() *gorm.DB {
	log := logging.MustGetLogger("log")

	log.Debug("db path: %s", viper.GetString("db.path"))
	log.Debug("db filename: %s", viper.GetString("db.filename"))

	db, err := gorm.Open(
		"sqlite3",
		filepath.Join(
			viper.GetString("db.path"),
			viper.GetString("db.filename"),
		),
	)
	if err != nil {
		log.Critical("Unable to open db file: %s", err)
		os.Exit(1)
	}
	if viper.GetString("logtype") == "debug" {
		db.LogMode(viper.GetBool("debug"))
	}
	db.CreateTable(new(Temperature))
	db.DB().Ping()

	return &db
}
