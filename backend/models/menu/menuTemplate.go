package menu

import (
	"time"

	"github.com/google/uuid"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/store"
)

type MenuTemplate struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	StoreTemplateID uuid.UUID `gorm:"type:uuid;not null"`
	StoreTemplate   store.StoreTemplate `gorm:"foreignKey:StoreTemplateID;constraint:OnDelete:CASCADE"`
	Name            string    `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (MenuTemplate) TableName() string {
	return "menu_templates"
}
