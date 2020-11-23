package main

import (
	"fmt"
	"github.com/Elephmoon/anagramDictionary/internal/backend/middleware"
	"github.com/Elephmoon/anagramDictionary/internal/backend/routes"
	"github.com/Elephmoon/anagramDictionary/internal/config"
	"github.com/Elephmoon/anagramDictionary/internal/migrations"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

func main() {
	conf := config.LoadConfig()
	logger := initLogger(conf.AppConfig)
	dbConn, err := initDB(conf.DBConfig)
	if err != nil {
		logger.Panic(err)
	}
	router := initAPI(conf.APIConfig, dbConn, logger)
	addr := fmt.Sprintf(":%s", conf.APIConfig.Port)
	logger.Infof("server start at %s", addr)
	logger.Panic(http.ListenAndServe(addr, router))
}

func initDB(conf config.DBConfig) (*gorm.DB, error) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Name)
	dbConn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't connect to database")
	}
	migrations.All(dbConn)
	return dbConn, nil
}

func initSwaggerSpec(conf config.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Host string
			Port string
		}{
			Host: conf.Host,
			Port: conf.Port,
		}
		tmpl := template.Must(template.ParseFiles("./api/api.yml"))
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func initLogger(conf config.AppConfig) logrus.FieldLogger {
	logger := logrus.New()
	if conf.Debug == config.On {
		logger.SetLevel(logrus.DebugLevel)
		logger.SetReportCaller(true)
	} else {
		logger.SetLevel(logrus.InfoLevel)
		logger.SetReportCaller(false)
	}
	logger.SetFormatter(&logrus.TextFormatter{})
	return logger.WithField("program", "anagramDictionary")
}

func initAPI(conf config.APIConfig, dbConn *gorm.DB, logger logrus.FieldLogger) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LogHTTPRequest(logger))
	swag := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./internal/static/swagger_dist")))
	router.PathPrefix("/swagger-ui/").Handler(swag)
	router.HandleFunc("/api.yml", initSwaggerSpec(conf))
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	routes.InitAPIRoutes(apiRouter, dbConn, logger)
	return router
}
