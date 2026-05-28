package menu

import (
	"github.com/google/uuid"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/product"
)

type MenuProduct struct {
	MenuID    uuid.UUID       `gorm:"type:uuid;primaryKey"`
	Menu      Menu            `gorm:"foreignKey:MenuID;constraint:OnDelete:CASCADE"`
	ProductID uuid.UUID       `gorm:"type:uuid;primaryKey"`
	Product   product.Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

func (MenuProduct) TableName() string {
	return "menu_products"
}
