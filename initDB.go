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

type Temperatures struct {
	Id   int `gorm:"primary_key"`
	Temp float64
	Date time.Time
}

type Eleves struct {
	Id 				int `gorm:"primary_key"`
	Identifiant		string
	MotDePasse		string
	Nom 			string
	Prenom 			string
	Email 			string
	Classe 			string
	Niveau 			string
	Sexe 			string
	DateNaissance 	time.Time
}

type Exams struct {
	Id int `gorm:"primary_key"`
	Date time.Time
	Score float64
}

type Questions struct {
	Id int `gorm:"primary_key"`
	Question string
}

type Reponses struct {
	Id int `gorm:"primary_key"`
	Reponse string
	IdQuestion int `sql:"type:bigint REFERENCES Questions(id)"`
}

type Solutions struct {
	Id int `gorm:"primary_key"`
	Solution string
	IdQuestion int `sql:"type:bigint REFERENCES Questions(id)"`
}

type Passer struct {
	Id int `gorm:"primary_key"`
	IdEleves int `sql:"type:bigint REFERENCES Eleves(id)"`
	IdExams int `sql:"type:bigint REFERENCES Exams(id)"`
}

type ContenirExamsQuestions struct {
	Id int `gorm:"primary_key"`
	IdExams int `sql:"type:bigint REFERENCES Exams(id)"`
	IdQuestions int `sql:"type:bigint REFERENCES Questions(id)"`
}

type ContenirExamsReponses struct {
	Id int `gorm:"primary_key"`
	IdExams int `sql:"type:bigint REFERENCES Exams(id)"`
	IdReponses int `sql:"type:bigint REFERENCES Reponses(id)"`
}

type Apporter struct {
	Id int `gorm:"primary_key"`
	IdQuestions int `sql:"type:bigint REFERENCES Questions(id)"`
	IdReponses int `sql:"type:bigint REFERENCES Reponses(id)"`
}

type Corriger struct {
	Id int `gorm:"primary_key"`
	IdSolution int `sql:"type:bigint REFERENCES Solutions(id)"`
	IdReponses int `sql:"type:bigint REFERENCES Reponses(id)"`
}

func Initdb() *gorm.DB {
	log := logging.MustGetLogger("log")

	log.Debug("db path: %s", viper.GetString("database.path"))
	log.Debug("db filename: %s", viper.GetString("database.filename"))

	db, err := gorm.Open(
		"sqlite3",
		filepath.Join(
			viper.GetString("database.path"),
			viper.GetString("database.filename"),
		),
	)

	log.Debug("db: %s", filepath.Join(viper.GetString("database.path"), viper.GetString("database.filename")))
	
	if err != nil {
		log.Critical("Unable to open db file: %s", err)
		os.Exit(1)
	}

	if viper.GetString("logtype") == "debug" {
		db.LogMode(viper.GetBool("debug"))
	}
	
	db.CreateTable(new(Temperatures))
	db.CreateTable(new(Eleves))
	db.CreateTable(new(Exams))
	db.CreateTable(new(Questions))
	db.CreateTable(new(Reponses))
	db.CreateTable(new(Solutions))
	db.CreateTable(new(Passer))
	db.CreateTable(new(ContenirExamsQuestions))
	db.CreateTable(new(ContenirExamsReponses))
	db.CreateTable(new(Apporter))
	db.CreateTable(new(Corriger))

	db.DB().Ping()

	return &db
}