package routes

import (
	wordhttp "github.com/Elephmoon/anagramDictionary/internal/backend/word/delivery/http"
	wordrepo "github.com/Elephmoon/anagramDictionary/internal/backend/word/repository"
	worducase "github.com/Elephmoon/anagramDictionary/internal/backend/word/usecase"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func InitAPIRoutes(router *mux.Router, dbConn *gorm.DB, logger logrus.FieldLogger) {
	wordRepo := wordrepo.NewDictionaryRepo(dbConn)
	wordUsecasee := worducase.NewWordUsecase(wordRepo)

	wordhttp.NewWordHandler(router, wordUsecasee, logger)
}
