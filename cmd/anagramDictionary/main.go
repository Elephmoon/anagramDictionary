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
	"html/template"
	"net/http"
)

func main() {
	conf := config.LoadConfig()
	dbConn, err := initDB(conf.DBConfig)
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	swag := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./internal/static/swagger_dist")))
	router.PathPrefix("/swagger-ui/").Handler(swag)
	router.HandleFunc("/api.yml", initSwaggerSpec(conf.APIConfig))
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	routes.InitApiRoutes(apiRouter, dbConn)
	addr := fmt.Sprintf(":%s", conf.APIConfig.Port)
	panic(http.ListenAndServe(addr, router))
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
			Host: "localhost",
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
