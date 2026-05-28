package models

import (
	"time"

	"github.com/google/uuid"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/codes"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/table"
)

type Order struct {
	ID           uuid.UUID          `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	SessionID    uuid.UUID          `json:"session_id" gorm:"type:uuid;not null"`
	TableSession table.TableSession `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE"`
	WaiterCodeID *uuid.UUID         `gorm:"type:uuid"`
	WaiterCode   *codes.WaiterCode  `gorm:"foreignKey:WaiterCodeID;constraint:OnDelete:SET NULL"`
	Status       string             `gorm:"type:order_status_enum;default:pending;not null"`
	Notes        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Order) TableName() string {
	return "orders"
}
