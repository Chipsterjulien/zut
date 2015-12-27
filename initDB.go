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
	// Username		string `sql:"not null;unique"`
	// Password		string `sql:"not null"`
	// Nom 			string `sql:"not null"`
	// Prenom 			string `sql:"not null"`
	// Email 			string `sql:"not null"`
	Username		string
	Password		string
	Nom 			string
	Prenom 			string
	Email 			string
	Classe 			string
	Sexe 			string
}

type Exams struct {
	Id int `gorm:"primary_key"`
	DebutExam time.Time
	FinExam time.Time
	NiveauExam string
	Fini bool
	Score float64
	IdEleve int `sql:"type:bigint REFERENCES Eleves(id)"`
}

type Questions struct {
	Id int `gorm:"primary_key"`
	Question string
	Type string
	Categorie string
}

type Reponses struct {
	Id int `gorm:"primary_key"`
	Reponse string
}

type Solutions struct {
	Id int `gorm:"primary_key"`
	Solution string
	IdQuestion int `sql:"type:bigint REFERENCES Questions(id)"`
}

type Propositions struct {
	Id int `gorm:"primary_key"`
	Proposition string
	IsMath bool
}

type ContenirExamsQuestions struct {
	Id int `gorm:"primary_key"`
	IdExams int `sql:"type:bigint REFERENCES Exams(id)"`
	IdQuestions int `sql:"type:bigint REFERENCES Questions(id)"`
	Done bool
}

type ContenirExamsReponses struct {
	Id int `gorm:"primary_key"`
	IdExams int `sql:"type:bigint REFERENCES Exams(id)"`
	IdReponses int `sql:"type:bigint REFERENCES Reponses(id)"`
}

type ContenirQuestionsPropositions struct {
	Id int `gorm:"primary_key"`
	IdQuestions int `sql:"type:bigint REFERENCES Questions(id)"`
	IdPropositions int `sql:"type:bigint REFERENCES Propositions(id)"`
}

type Apporter struct {
	Id int `gorm:"primary_key"`
	IdQuestions int `sql:"type:bigint REFERENCES Questions(id)"`
	IdReponses int `sql:"type:bigint REFERENCES Reponses(id)"`
}

type Corriger struct {
	Id int `gorm:"primary_key"`
	IdReponses int `sql:"type:bigint REFERENCES Reponses(id)"`
	IdSolution int `sql:"type:bigint REFERENCES Solutions(id)"`
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
	db.CreateTable(new(Propositions))
	db.CreateTable(new(ContenirExamsQuestions))
	db.CreateTable(new(ContenirExamsReponses))
	db.CreateTable(new(ContenirQuestionsPropositions))
	db.CreateTable(new(Apporter))
	db.CreateTable(new(Corriger))

	db.DB().Ping()

	return &db
}