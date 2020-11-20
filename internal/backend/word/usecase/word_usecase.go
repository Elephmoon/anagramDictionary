package usecase

import (
	"github.com/Elephmoon/anagramDictionary/internal/backend/word"
	"github.com/Elephmoon/anagramDictionary/internal/helpers"
	"github.com/Elephmoon/anagramDictionary/internal/models"
	"strings"
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
	return wu.wordRepo.DeleteWord(word)
}

func (wu *wordUsecase) AddWords(words []string) error {
	wrds := make([]models.Word, len(words))
	for i, wrd := range words {
		wrds[i].Word = words[i]
		wrds[i].SortedWord = helpers.SortWord(strings.ToLower(wrd))
	}
	return wu.wordRepo.AddDictionary(wrds)
}
