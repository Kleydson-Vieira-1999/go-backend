package table

import (
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&Table{},
		&TableSession{},
	)
}
