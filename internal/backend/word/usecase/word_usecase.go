package usecase

import (
	"errors"
	"github.com/Elephmoon/anagramDictionary/internal/backend/word"
	"github.com/Elephmoon/anagramDictionary/internal/helpers"
	"github.com/Elephmoon/anagramDictionary/internal/models"
	"github.com/go-playground/validator/v10"
	"strings"
)

var errWordEmpty = errors.New("word cannot be empty")

type wordUsecase struct {
	wordRepo word.Repository
}

func NewWordUsecase(wordRepo word.Repository) word.Usecase {
	return &wordUsecase{
		wordRepo: wordRepo,
	}
}

func (wu *wordUsecase) ShowDictionary(offset, limit string) ([]*models.Word, error) {
	queryParams, err := helpers.GetQueryParams(offset, limit)
	if err != nil {
		return nil, err
	}
	return wu.wordRepo.GetDictionary(queryParams.Offset, queryParams.Limit)
}

func (wu *wordUsecase) DeleteWord(word string) error {
	if word == "" {
		return errWordEmpty
	}
	return wu.wordRepo.DeleteWord(word)
}

func (wu *wordUsecase) AddWords(words *models.CreateReq) error {
	err := validateCreateReq(words)
	if err != nil {
		return err
	}
	wrds := make([]models.Word, len(words.Words))
	for i, wrd := range words.Words {
		wrds[i].Word = words.Words[i]
		wrds[i].SortedWord = helpers.SortWord(strings.ToLower(wrd))
	}
	return wu.wordRepo.AddDictionary(wrds)
}

func validateCreateReq(words *models.CreateReq) error {
	validate := validator.New()
	return validate.Struct(words)
}
