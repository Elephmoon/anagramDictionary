package http

import (
	"encoding/json"
	"github.com/Elephmoon/anagramDictionary/internal/backend/word"
	"github.com/Elephmoon/anagramDictionary/internal/helpers"
	"github.com/Elephmoon/anagramDictionary/internal/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type WordHandler struct {
	WordUsecase word.Usecase
	Logger      logrus.FieldLogger
}

func NewWordHandler(router *mux.Router, wordUsecase word.Usecase, logger logrus.FieldLogger) {
	handler := WordHandler{
		WordUsecase: wordUsecase,
		Logger:      logger,
	}
	wordSubRouter := router.PathPrefix("/words").Subrouter()
	wordSubRouter.HandleFunc("", handler.get).Methods("GET")
	wordSubRouter.HandleFunc("", handler.delete).Methods("DELETE")
	wordSubRouter.HandleFunc("", handler.addWords).Methods("POST")
	wordSubRouter.HandleFunc("/anagrams", handler.searchAnagram).Methods("GET")
}

func (wh *WordHandler) get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	offset := queryParams.Get("offset")
	limit := queryParams.Get("limit")

	data, err := wh.WordUsecase.ShowDictionary(offset, limit)
	if err != nil {
		if err := helpers.GenerateHTTPErrorResp(w, err); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			helpers.LogHTTPError(wh.Logger, r, err)
			return
		}
		helpers.LogHTTPError(wh.Logger, r, err)
		return
	}
	w.Header().Add(helpers.KeyContentType, helpers.JSONContentType)
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		helpers.LogHTTPError(wh.Logger, r, err)
		return
	}
}

func (wh *WordHandler) delete(w http.ResponseWriter, r *http.Request) {
	err := wh.WordUsecase.DeleteWord(r.URL.Query().Get("word"))
	if err != nil {
		if err := helpers.GenerateHTTPErrorResp(w, err); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			helpers.LogHTTPError(wh.Logger, r, err)
			return
		}
		helpers.LogHTTPError(wh.Logger, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (wh *WordHandler) addWords(w http.ResponseWriter, r *http.Request) {
	createReq := &models.CreateReq{}
	err := json.NewDecoder(r.Body).Decode(createReq)
	if err != nil {
		if err := helpers.GenerateHTTPErrorResp(w, err); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			helpers.LogHTTPError(wh.Logger, r, err)
			return
		}
		helpers.LogHTTPError(wh.Logger, r, err)
		return
	}
	err = wh.WordUsecase.AddWords(createReq)
	if err != nil {
		if err := helpers.GenerateHTTPErrorResp(w, err); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			helpers.LogHTTPError(wh.Logger, r, err)
			return
		}
		helpers.LogHTTPError(wh.Logger, r, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (wh *WordHandler) searchAnagram(w http.ResponseWriter, r *http.Request) {
	anagramResponse, err := wh.WordUsecase.AnagramSearch(r.URL.Query().Get("word"))
	if err != nil {
		if err := helpers.GenerateHTTPErrorResp(w, err); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			helpers.LogHTTPError(wh.Logger, r, err)
			return
		}
		helpers.LogHTTPError(wh.Logger, r, err)
		return
	}
	w.Header().Add(helpers.KeyContentType, helpers.JSONContentType)
	err = json.NewEncoder(w).Encode(anagramResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		helpers.LogHTTPError(wh.Logger, r, err)
		return
	}
}
