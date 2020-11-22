package repository

import (
	"fmt"
	"github.com/Elephmoon/anagramDictionary/internal/backend/word"
	"github.com/Elephmoon/anagramDictionary/internal/helpers"
	"github.com/Elephmoon/anagramDictionary/internal/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type wordRepository struct {
	DBConn    *gorm.DB
	TableName string
}

func NewDictionaryRepo(dbConn *gorm.DB) word.Repository {
	return &wordRepository{
		DBConn:    dbConn,
		TableName: dbConn.NewScope(models.Word{}).TableName(),
	}
}

func (wr *wordRepository) GetDictionary(offset, limit int) ([]*models.Word, error) {
	dict := make([]*models.Word, 0)
	err := wr.DBConn.Offset(offset).Limit(limit).Find(&dict).Error
	if err != nil {
		return nil, wr.wrapError(err)
	}
	return dict, nil
}

func (wr *wordRepository) AddDictionary(words []models.Word) error {
	for _, wrd := range words {
		err := wr.DBConn.Create(&wrd).Error
		if err != nil {
			return wr.wrapError(errors.Wrap(err, fmt.Sprintf("word: %s", wrd.Word)))
		}
	}
	return nil
}

func (wr *wordRepository) DeleteWord(word string) error {
	wrd := models.Word{}
	rowsAffected := wr.DBConn.Delete(wrd, "word = ?", word).RowsAffected
	if rowsAffected == 0 {
		return wr.wrapError(helpers.ErrRecordNotFound)
	}
	return nil
}

func (wr *wordRepository) AnagramSearch(sortedWord string) ([]*models.Word, error) {
	wrds := make([]*models.Word, 0)
	err := wr.DBConn.Where("sorted_word = ?", sortedWord).Find(wrds).Error
	if err != nil {
		return nil, wr.wrapError(err)
	}
	return wrds, nil
}

func (wr *wordRepository) wrapError(err error) error {
	return errors.Wrap(err, wr.TableName)
}
