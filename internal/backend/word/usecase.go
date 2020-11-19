package word

type Usecase interface {
	ShowDictionary(offset, limit int) ([]string, error)
	DeleteWord(word string) error
	AddWords(words []string) error
}
