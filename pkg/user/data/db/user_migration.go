package db

import (
	"github.com/atefeh-syf/yumigo/pkg/user/data/models"
	"gorm.io/gorm"
)

// Migrate Entity Here
func MigrateEntities(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.UserRole{},
	)
}
