package db

import (
	"github.com/atefeh-syf/E-Wallet/data/models"
	"gorm.io/gorm"
)

// Migrate Entity Here
func MigrateEntities(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Wallet{},
		&models.Transaction{},
	)
}
