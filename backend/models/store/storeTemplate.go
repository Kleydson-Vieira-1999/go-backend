package store

import (
	"time"

	"github.com/google/uuid"
)

type StoreTemplate struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string    `gorm:"not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (StoreTemplate) TableName() string {
	return "store_templates"
}
