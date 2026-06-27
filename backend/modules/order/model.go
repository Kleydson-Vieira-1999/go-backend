package order

import (
	"time"

	"github.com/google/uuid"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/product"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/table"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/waiter"
)

type Order struct {
	ID           uuid.UUID          `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	SessionID    uuid.UUID          `json:"session_id" gorm:"type:uuid;not null"`
	TableSession table.TableSession `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE"`
	WaiterCodeID *uuid.UUID         `gorm:"type:uuid"`
	WaiterCode   *waiter.WaiterCode `gorm:"foreignKey:WaiterCodeID;constraint:OnDelete:SET NULL"`
	Status       string             `gorm:"type:order_status_enum;default:pending;not null"`
	Notes        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Order) TableName() string {
	return "orders"
}

type OrderItem struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	OrderID   uuid.UUID       `gorm:"type:uuid;not null"`
	Order     Order           `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	ProductID uuid.UUID       `gorm:"type:uuid;not null"`
	Product   product.Product `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT"`
	Quantity  int             `gorm:"not null;check:quantity > 0"`
	UnitCost  int             `gorm:"not null"`
	UnitPrice int             `gorm:"not null"`
	Notes     string
	UpdatedAt string
}

func (OrderItem) TableName() string {
	return "order_items"
}
