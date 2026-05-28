package codes

import (
	"github.com/google/uuid"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/store"
)

type KitchenCode struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	StoreID  uuid.UUID `gorm:"type:uuid;not null"`
	Store    store.Store     `gorm:"foreignKey:StoreID;constraint:OnDelete:CASCADE"`
	Code     string    `gorm:"not null;uniqueIndex:idx_kitchen_codes"`
	Label    string
	IsActive bool      `gorm:"default:true"`
	UpdatedAt string
}

func (KitchenCode) TableName() string {
	return "kitchen_codes"
}
