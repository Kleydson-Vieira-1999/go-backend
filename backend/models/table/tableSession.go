package table

import (
	"time"

	"github.com/google/uuid"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/store"
)

type TableSession struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	StoreID   uuid.UUID `gorm:"type:uuid;not null;index:idx_table_sessions_status"`
	Store     store.Store     `gorm:"foreignKey:StoreID;constraint:OnDelete:CASCADE"`
	TableID   uuid.UUID `gorm:"type:uuid;not null"`
	Table     Table     `gorm:"foreignKey:TableID;constraint:OnDelete:RESTRICT"`
	Status    string    `gorm:"type:session_status_enum;default:active;not null;index:idx_table_sessions_status"`
	OpenedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ClosedAt  *time.Time
	UpdatedAt time.Time
}

func (TableSession) TableName() string {
	return "table_sessions"
}
