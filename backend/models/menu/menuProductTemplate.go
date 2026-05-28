package menu

import (
	"github.com/google/uuid"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/product"
)

type MenuProductTemplate struct {
	MenuTemplateID    uuid.UUID               `gorm:"type:uuid;primaryKey"`
	MenuTemplate      MenuTemplate            `gorm:"foreignKey:MenuTemplateID;constraint:OnDelete:CASCADE"`
	ProductTemplateID uuid.UUID               `gorm:"type:uuid;primaryKey"`
	ProductTemplate   product.ProductTemplate `gorm:"foreignKey:ProductTemplateID;constraint:OnDelete:CASCADE"`
}

func (MenuProductTemplate) TableName() string {
	return "menu_product_templates"
}
