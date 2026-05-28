package store

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID          uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	StoreTemplateID uuid.UUID `json:"store_template_id" gorm:"type:uuid"`
	Name            string    `json:"name" gorm:"not null"`
	Picture         string    `json:"picture"`
	Type            string    `json:"type" `
	Description     string    `json:"description"`
	IsActive        bool      `json:"is_active" gorm:"default:true"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (Store) TableName() string {
	return "stores"
}
