package word

import "github.com/Elephmoon/anagramDictionary/internal/models"

type Usecase interface {
	ShowDictionary(offset, limit string) ([]*models.Word, error)
	DeleteWord(word string) error
	AddWords(words *models.CreateReq) error
	AnagramSearch(word string) (models.AnagramResponse, error)
}
