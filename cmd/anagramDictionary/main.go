package main

import (
	"fmt"
	"github.com/Elephmoon/anagramDictionary/internal/backend/routes"
	"github.com/Elephmoon/anagramDictionary/internal/config"
	"github.com/Elephmoon/anagramDictionary/internal/migrations"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	"net/http"
)

func main() {
	conf := config.LoadConfig()
	dbConn, err := initDB(conf.DBConfig)
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	routes.InitApiRoutes(apiRouter, dbConn)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.APIConfig.Port),
		Handler: apiRouter,
	}
	panic(server.ListenAndServe())
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
