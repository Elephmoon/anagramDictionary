package usecase

import (
	"errors"
	"fmt"
	"github.com/Elephmoon/anagramDictionary/internal/backend/word"
	"github.com/Elephmoon/anagramDictionary/internal/helpers"
	"github.com/Elephmoon/anagramDictionary/internal/models"
	"github.com/go-playground/validator/v10"
	"strings"
)

const (
	maxWordsLen        = 100
	errExceedLenFormat = "max length words is %d"
)

var (
	errWordEmpty = errors.New("word cannot be empty")
)

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
		wrds[i].SortedWord, err = helpers.SortWord(strings.ToLower(wrd))
		if err != nil {
			return err
		}
	}
	return wu.wordRepo.AddDictionary(wrds)
}

func (wu *wordUsecase) AnagramSearch(word string) (models.AnagramResponse, error) {
	var answer models.AnagramResponse
	if word == "" {
		return answer, errWordEmpty
	}
	sortedWord, err := helpers.SortWord(word)
	if err != nil {
		return answer, err
	}
	words, err := wu.wordRepo.AnagramSearch(sortedWord)
	if err != nil {
		return answer, err
	}
	return generateAnagramResponse(word, words), nil
}

func generateAnagramResponse(searchWord string, words []*models.Word) models.AnagramResponse {
	anagrams := make([]string, len(words))
	for i, wrd := range words {
		anagrams[i] = wrd.Word
	}
	return models.AnagramResponse{
		Word:     searchWord,
		Anagrams: anagrams,
	}
}

func validateCreateReq(words *models.CreateReq) error {
	validate := validator.New()
	err := validate.Struct(words)
	if err != nil {
		return err
	}
	if len(words.Words) > maxWordsLen {
		return fmt.Errorf(errExceedLenFormat, maxWordsLen)
	}
	return nil
}
