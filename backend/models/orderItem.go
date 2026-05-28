package models

import (
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/product"
	"github.com/google/uuid"
)

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
