package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Elephmoon/anagramDictionary/internal/backend/word/mocks"
	"github.com/Elephmoon/anagramDictionary/internal/models"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWordHandler_get(t *testing.T) {
	type fields struct {
		WordUsecase *mocks.Usecase
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	type testData struct {
		name   string
		fields fields
		args   args
	}
	test1 := func() testData {
		mockDictionary := make([]*models.Word, 0)
		var mockWord models.Word
		err := faker.FakeData(&mockWord)
		assert.NoError(t, err)

		mockDictionary = append(mockDictionary, &mockWord)
		mockUsecase := new(mocks.Usecase)
		mockUsecase.On("ShowDictionary", "", "").Return(mockDictionary, nil)

		req, err := http.NewRequest("GET", "words", nil)
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		return testData{
			name: "no error",
			fields: fields{
				WordUsecase: mockUsecase,
			},
			args: args{
				w: rec,
				r: req,
			},
		}
	}
	test2 := func() testData {
		mockDictionary := make([]*models.Word, 0)
		var mockWord models.Word
		err := faker.FakeData(&mockWord)
		assert.NoError(t, err)

		mockDictionary = append(mockDictionary, &mockWord)
		mockUsecase := new(mocks.Usecase)
		offset := "asdas"
		mockUsecase.On("ShowDictionary", offset, "").Return(mockDictionary, nil)
		req, err := http.NewRequest("GET", fmt.Sprintf("/words/?offset=%s", offset), nil)
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		return testData{
			name: "wrong offset parameters",
			fields: fields{
				WordUsecase: mockUsecase,
			},
			args: args{
				w: rec,
				r: req,
			},
		}
	}
	test3 := func() testData {
		mockDictionary := make([]*models.Word, 0)
		var mockWord models.Word
		err := faker.FakeData(&mockWord)
		assert.NoError(t, err)

		mockDictionary = append(mockDictionary, &mockWord)
		mockUsecase := new(mocks.Usecase)
		limit := "asdas"
		mockUsecase.On("ShowDictionary", "", limit).Return(mockDictionary, nil)
		req, err := http.NewRequest("GET", fmt.Sprintf("/words/?limit=%s", limit), nil)
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		return testData{
			name: "wrong limit parameters",
			fields: fields{
				WordUsecase: mockUsecase,
			},
			args: args{
				w: rec,
				r: req,
			},
		}
	}
	tests := []testData{test1(), test2(), test3()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WordHandler{
				WordUsecase: tt.fields.WordUsecase,
			}
			handler := http.HandlerFunc(w.get)
			handler.ServeHTTP(tt.args.w, tt.args.r)

			assert.Equal(t, http.StatusOK, tt.args.w.Code)
			tt.fields.WordUsecase.AssertExpectations(t)
		})
	}
}

func TestWordHandler_delete(t *testing.T) {
	type fields struct {
		WordUsecase *mocks.Usecase
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	type testData struct {
		name       string
		fields     fields
		args       args
		wantStatus int
	}
	test1 := func() testData {
		var mockWord models.Word
		err := faker.FakeData(&mockWord)
		assert.NoError(t, err)

		deletedWord := mockWord.Word

		mockUsecase := new(mocks.Usecase)
		mockUsecase.On("DeleteWord", deletedWord).Return(nil)

		req, err := http.NewRequest("DELETE", fmt.Sprintf("/words/?word=%s", deletedWord), nil)
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		return testData{
			name: "no error",
			fields: fields{
				WordUsecase: mockUsecase,
			},
			args: args{
				w: rec,
				r: req,
			},
			wantStatus: http.StatusNoContent,
		}
	}
	test2 := func() testData {
		var mockWord models.Word
		err := faker.FakeData(&mockWord)
		assert.NoError(t, err)

		deletedWord := mockWord.Word

		mockUsecase := new(mocks.Usecase)
		mockUsecase.On("DeleteWord", deletedWord).Return(errors.New("word not found"))

		req, err := http.NewRequest("DELETE", fmt.Sprintf("/words/?word=%s", deletedWord), nil)
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		return testData{
			name: "not found word",
			fields: fields{
				WordUsecase: mockUsecase,
			},
			args: args{
				w: rec,
				r: req,
			},
			wantStatus: http.StatusBadRequest,
		}
	}
	tests := []testData{test1(), test2()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WordHandler{
				WordUsecase: tt.fields.WordUsecase,
			}
			handler := http.HandlerFunc(w.delete)
			handler.ServeHTTP(tt.args.w, tt.args.r)

			assert.Equal(t, tt.wantStatus, tt.args.w.Code)
			tt.fields.WordUsecase.AssertExpectations(t)
		})
	}
}

func TestWordHandler_addWords(t *testing.T) {
	type fields struct {
		WordUsecase *mocks.Usecase
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	type testData struct {
		name       string
		fields     fields
		args       args
		wantStatus int
	}
	test1 := func() testData {
		var mockReq models.CreateReq
		err := faker.FakeData(&mockReq)
		assert.NoError(t, err)
		jsonData, err := json.Marshal(mockReq)
		assert.NoError(t, err)

		mockUsecase := new(mocks.Usecase)
		mockUsecase.On("AddWords", mock.AnythingOfType("*models.CreateReq")).Return(nil)

		req, err := http.NewRequest("POST", "/words", strings.NewReader(string(jsonData)))
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		return testData{
			name: "no error",
			fields: fields{
				WordUsecase: mockUsecase,
			},
			args: args{
				w: rec,
				r: req,
			},
			wantStatus: http.StatusCreated,
		}
	}
	test2 := func() testData {
		var mockReq models.CreateReq
		err := faker.FakeData(&mockReq)
		assert.NoError(t, err)
		jsonData, err := json.Marshal(mockReq)
		assert.NoError(t, err)

		mockUsecase := new(mocks.Usecase)
		mockUsecase.On("AddWords", mock.AnythingOfType("*models.CreateReq")).Return(errors.New("bad req"))

		req, err := http.NewRequest("POST", "/words", strings.NewReader(string(jsonData)))
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		return testData{
			name: "bad request",
			fields: fields{
				WordUsecase: mockUsecase,
			},
			args: args{
				w: rec,
				r: req,
			},
			wantStatus: http.StatusBadRequest,
		}
	}
	tests := []testData{test1(), test2()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WordHandler{
				WordUsecase: tt.fields.WordUsecase,
			}
			handler := http.HandlerFunc(w.addWords)
			handler.ServeHTTP(tt.args.w, tt.args.r)

			assert.Equal(t, tt.wantStatus, tt.args.w.Code)
			tt.fields.WordUsecase.AssertExpectations(t)
		})
	}
}

func TestWordHandler_searchAnagram(t *testing.T) {
	type fields struct {
		WordUsecase *mocks.Usecase
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	type testData struct {
		name       string
		fields     fields
		args       args
		wantStatus int
	}
	test1 := func() testData {
		var mockResp models.AnagramResponse
		err := faker.FakeData(&mockResp)
		assert.NoError(t, err)

		mockUsecase := new(mocks.Usecase)
		mockUsecase.On("AnagramSearch", mockResp.Word).Return(mockResp, nil)

		req, err := http.NewRequest("GET", fmt.Sprintf("/words/?word=%s", mockResp.Word), nil)
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		return testData{
			name: "no error",
			fields: fields{
				WordUsecase: mockUsecase,
			},
			args: args{
				w: rec,
				r: req,
			},
			wantStatus: http.StatusOK,
		}
	}
	test2 := func() testData {
		var word string
		err := faker.FakeData(&word)
		assert.NoError(t, err)

		mockUsecase := new(mocks.Usecase)
		mockUsecase.On("AnagramSearch", word).Return(models.AnagramResponse{}, errors.New("bad request"))

		req, err := http.NewRequest("GET", fmt.Sprintf("/words/?word=%s", word), nil)
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		return testData{
			name: "bad request",
			fields: fields{
				WordUsecase: mockUsecase,
			},
			args: args{
				w: rec,
				r: req,
			},
			wantStatus: http.StatusBadRequest,
		}
	}
	tests := []testData{test1(), test2()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WordHandler{
				WordUsecase: tt.fields.WordUsecase,
			}
			handler := http.HandlerFunc(w.searchAnagram)
			handler.ServeHTTP(tt.args.w, tt.args.r)

			assert.Equal(t, tt.wantStatus, tt.args.w.Code)
			tt.fields.WordUsecase.AssertExpectations(t)
		})
	}
}
