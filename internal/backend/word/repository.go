package word

import "github.com/Elephmoon/anagramDictionary/internal/models"

type Repository interface {
	GetDictionary(offset, limit int) ([]*models.Word, error)
	AddDictionary(words []models.Word) error
	DeleteWord(word string) error
	AnagramSearch(sortedWord string) ([]*models.Word, error)
}
