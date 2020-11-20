package migrations

import (
	"github.com/Elephmoon/anagramDictionary/internal/models"
	"github.com/jinzhu/gorm"
)

func All(dbConn *gorm.DB) {
	dbConn.Debug().AutoMigrate(
		&models.Word{},
	)
}
