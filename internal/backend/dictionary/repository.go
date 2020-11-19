package dictionary

type Repository interface {
	GetDictionary(offset, limit int) ([]string, error)
	AddWords(words []string) error
	DeleteWord(word string) error
}
