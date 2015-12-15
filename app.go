package main

import (
	"strconv"
	"time"
	"math/rand"

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

// Init DB
func NewRessource(db *gorm.DB) Ressource {
	return Ressource{
		db: db,
	}
}

func authUserKeyNotFound(c *gin.Context) {
	c.Header("WWW-Authenticate", "Basic realm=Authorization Required")
	c.AbortWithStatus(401)
}

func (r *Ressource) CreateNewExam(c *gin.Context) {
	log := logging.MustGetLogger("log")

	user, found := c.Get(basicAuthWithDBForGin.AuthUserKey)
	if !found {
		authUserKeyNotFound(c)
	} else {
		log.Debug("Utilisateur: %s", user)
		niveau := c.Param("niveau")
		log.Debug("Niveau: %s", niveau)
		
		// Obtenir l'id de l'user
		var eleve = Eleves{}
		if err := r.db.Select("eleves.id").Table("eleves").Where("eleves.username=?", user).Find(&eleve).Error; err != nil {
			log.Warning("Erreur: %v", err)
			// Retourner un code d'erreur à l'application client (erreur 500 par exemple)
			return
		}
		log.Debug("Id: %v", eleve.Id)

		// Créer un exam
		// insert into exams (debut_exam, niveau_exam, id_eleve) values ("", "", "");
		exam := Exams{
			DebutExam: time.Now(),
			NiveauExam: c.Param("niveau"),
			IdEleve: eleve.Id,
		}

		if err := r.db.Debug().Save(&exam).Error; err != nil {
			log.Debug("Erreur: %v", err)
			// Retourner un code d'erreur à l'application client (erreur 500 par exemple)
			return
		}

		log.Debug("Exam après enregistrement: %v", exam)

		// Obtenir la liste des questions
		// select * from questions
		var questions = []Questions{}
		if err := r.db.Find(&questions).Error; err != nil {
			log.Critical("Erreur: %v", err)
			// Retourner un code d'erreur à l'application client (erreur 500 par exemple)
			return
		}
		log.Debug("Questions: %v", questions)

		// Sélectionner 50 questions
		// Tant que la taille de la liste est supérieure à x, on supprime des éléments aléatoires

		// On passe les id sur une map pour une suppression plus facile
		questionsIdMap := make(map[int]int)
		for num, q := range questions {
			questionsIdMap[num] = q.Id
		}

		for len(questionsIdMap) > 5 {
			num := rand.Intn(len(questionsIdMap))
			delete(questionsIdMap, num)
		}

		// On assigne les questions à l'exam
		for _, id := range questionsIdMap {
			examsQuestions := ContenirExamsQuestions{
				IdExams: exam.Id,
				IdQuestions: id,
			}
			if err := r.db.Debug().Save(&examsQuestions).Error; err != nil {
				log.Critical("Erreur: %v", err)
				// Retourner un code d'erreur à l'application client (erreur 500 par exemple)
				return
			}
		}
		// Une fois créer, envoyer un code 200 avec le num de l'exam
	}
}

func (r *Ressource) ListOfFinishedExams(c *gin.Context) {
	log := logging.MustGetLogger("log")

	user, found := c.Get(basicAuthWithDBForGin.AuthUserKey)
	if !found {
		authUserKeyNotFound(c)
	} else {
		log.Debug("User: %s", user)
		exams := []Exams{}

		if err := r.db.Select("exams.id, exams.debut_exam, exams.score").Table("exams").Joins("INNER JOIN eleves ON exams.id_eleve=eleves.id").Where("eleves.username=? and exams.fini=?", user, "true").Find(&exams).Error; err != nil {
			log.Debug("Impossible d'exécuter la recherche: 'Trouver la liste des exams non finis de %s'", user)
			c.JSON(500, nil)
		} else {
			log.Debug("Taille: %v | Exams: %v", len(exams), exams)
			c.JSON(200, exams)
		}
	}
}

func (r *Ressource) ListOfUnfinishedExams(c *gin.Context) {
	log := logging.MustGetLogger("log")

	user, ok := c.Get(basicAuthWithDBForGin.AuthUserKey)
	if !ok {
		authUserKeyNotFound(c)
	} else {
		log.Debug("User: %s", user)
		exams := []Exams{}

		if err := r.db.Select("exams.id, exams.debut_exam").Table("exams").Joins("INNER JOIN eleves ON exams.id_eleve=eleves.id").Where("eleves.username=? and exams.fini=?", user, "false").Find(&exams).Error; err != nil {
			log.Debug("Impossible d'exécuter la recherche: 'Trouver la liste des exams non finis de %s'", user)
			c.JSON(500, nil)
		} else {
			log.Debug("Taille: %v | Exams: %v", len(exams), exams)
			c.JSON(200, exams)
		}
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
		authorized.GET("listOfFinishedExams", r.ListOfFinishedExams)
		authorized.GET("createNewExam/:niveau", r.CreateNewExam)
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
