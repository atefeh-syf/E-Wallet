package db

import (
	"github.com/atefeh-syf/yumigo/pkg/wallet/data/models"
	"gorm.io/gorm"
)

// Migrate Entity Here
func MigrateEntities(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Wallet{},
		&models.Transaction{},
	)
}
