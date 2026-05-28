package table

import (
	"github.com/google/uuid"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/store"
)

type Table struct {
	ID         uuid.UUID   `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	StoreID    uuid.UUID   `gorm:"type:uuid;not null"`
	Store      store.Store `gorm:"foreignKey:StoreID;constraint:OnDelete:CASCADE"`
	Identifier string      `gorm:"not null;uniqueIndex:idx_table_identifier"`
	IsActive   bool        `gorm:"default:true"`
	UpdatedAt  string
}

func (Table) TableName() string {
	return "tables"
}
