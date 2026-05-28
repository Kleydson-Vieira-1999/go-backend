package menu

import (
	"github.com/google/uuid"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/store"
)

type Menu struct {
	ID        uuid.UUID   `json:"id"        gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	StoreID   uuid.UUID   `json:"store_id"  gorm:"type:uuid;not null"`
	Store     store.Store `json:"store"     gorm:"foreignKey:StoreID;constraint:OnDelete:CASCADE"`
	Name      string      `json:"name"      gorm:"not null"`
	IsActive  bool        `json:"is_active" gorm:"default:true"`
	UpdatedAt string      `json:"updated_at" `
}

func (Menu) TableName() string {
	return "menus"
}
