package store

import (
	"time"

	"github.com/google/uuid"
)

type StoreBalance struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	StoreID        uuid.UUID `gorm:"type:uuid;not null;unique"`
	Store          Store     `gorm:"foreignKey:StoreID;constraint:OnDelete:CASCADE"`
	CurrentBalance int       `gorm:"default:0"`
	TotalProfit    int       `gorm:"default:0"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

func (StoreBalance) TableName() string {
	return "store_balance"
}
