package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"github.com/chipsterjulien/basicAuthWithDBForGin"
)

// Struct permit to have DB access on method
type Ressource struct {
	db *gorm.DB
}

type ExamsInfo struct {
	id int
	debut string
}

// Init DB
func NewRessource(db *gorm.DB) Ressource {
	return Ressource{
		db: db,
	}
}

func (r *Ressource) ListOfUnfinishedExams(c *gin.Context) {
	log := logging.MustGetLogger("log")
	log.Debug("-------------------")
	log.Debug("Fonction ListOfExem")
	log.Debug("-------------------")
	user, ok := c.Get(basicAuthWithDBForGin.AuthUserKey)
	if !ok {
		return
	} else {
		log.Debug("User: %s", user)
		
		rows, err := r.db.Debug().Raw("SELECT exams.id, exams.debut_exam FROM exams INNER JOIN eleves ON exams.id_eleve = eleves.id WHERE (eleves.username = ? and exams.fini = ?)", user, false).Rows()
		if err != nil {
			log.Debug("%v", err)
		}
		defer rows.Close()

		var id int
		var debut string

		for rows.Next() {
			rows.Scan(&id, &debut)
			log.Debug("id: %d | debut: %v", id, debut)
		}

		exams := []ExamsInfo{}

		r.db.Debug().Select("exams.id, exams.debut_exam").Table("exams").Joins("INNER JOIN eleves ON exams.id_eleve=eleves.id").Where("eleves.username = ? and exams.fini = ?", user, false).Find(&exams)
		log.Debug("Exams: %v", exams)

		// query := r.db.Query("SELECT * FROM users WHERE username=$1", user)
		// Je veux la liste des exams

		// SELECT ex.id, ex.debut_exam FROM exams as ex INNER JOIN eleves as el ON ex.id_eleve=el.id WHERE el.username="julien" and ex.fini="false"

		// SELECT exams.id exams.debut_exam FROM exams INNER JOIN eleves ON exams.id_eleve = eleves.id WHERE eleves.username = 'julien' and exams.fini = 'false';
		// SELECT ex.id, ex.debut_exam FROM exams as ex INNER JOIN eleves as el ON ex.id_eleve=el.id WHERE el.username="julien" and ex.fini="false";

		// log.Debug("Retour: %v", query)
		// c.JSON(200, query)
	}
}

func startApp(db *gorm.DB) {
	log := logging.MustGetLogger("log")

	if viper.GetString("logtype") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.Default()
	r := NewRessource(db)

	g.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	g.Static("/static", "./static")

	authorized := g.Group("authorized", basicAuthWithDBForGin.BasicAuthWithDB(db, "Eleves"))
	{
		// Possibilité de supprimer la ligne suivante en ne modifiant que les headers et en allant à la page suivante
		authorized.GET("/", func(c *gin.Context){})
		authorized.GET("listOfUnfinishedExams", r.ListOfUnfinishedExams)
	}

	log.Debug("Port: %d", viper.GetInt("server.port"))
	g.Run(":" + strconv.Itoa(viper.GetInt("server.port")))
}

func main() {
	confPath := "cfg"
	confFilename := "server"
	logFilename := "error.log"

	fd := initLogging(&logFilename)
	defer fd.Close()

	loadConfig(&confPath, &confFilename)
	dbmap := Initdb()

	startApp(dbmap)
}
