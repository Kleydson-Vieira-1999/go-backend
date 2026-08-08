package menu

import (
	"time"

	"github.com/google/uuid"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/product"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/store"
)

type Menu struct {
	ID        uuid.UUID   `json:"id"        gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	StoreID   uuid.UUID   `json:"store_id"  gorm:"type:uuid;not null"`
	Store     store.Store `json:"store"     gorm:"foreignKey:StoreID;constraint:OnDelete:CASCADE"`
	Code      string      `json:"menu_code" gorn:""`
	Name      string      `json:"name"      gorm:"not null"`
	IsActive  bool        `json:"is_active" gorm:"default:true"`
	UpdatedAt string      `json:"updated_at" `
}

func (Menu) TableName() string {
	return "menus"
}

type MenuProduct struct {
	MenuID      uuid.UUID       `gorm:"type:uuid;primaryKey"`
	Menu        Menu            `gorm:"foreignKey:MenuID;constraint:OnDelete:CASCADE"`
	ProductID   uuid.UUID       `gorm:"type:uuid;primaryKey"`
	Product     product.Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	IsAvailable bool            `json:"is_available" gorm:"default:true"`
}

func (MenuProduct) TableName() string {
	return "menu_products"
}

type MenuProductTemplate struct {
	MenuTemplateID    uuid.UUID               `gorm:"type:uuid;primaryKey"`
	MenuTemplate      MenuTemplate            `gorm:"foreignKey:MenuTemplateID;constraint:OnDelete:CASCADE"`
	ProductTemplateID uuid.UUID               `gorm:"type:uuid;primaryKey"`
	ProductTemplate   product.ProductTemplate `gorm:"foreignKey:ProductTemplateID;constraint:OnDelete:CASCADE"`
}

func (MenuProductTemplate) TableName() string {
	return "menu_product_templates"
}

type MenuTemplate struct {
	ID              uuid.UUID           `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	StoreTemplateID uuid.UUID           `gorm:"type:uuid;not null"`
	StoreTemplate   store.StoreTemplate `gorm:"foreignKey:StoreTemplateID;constraint:OnDelete:CASCADE"`
	Name            string              `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (MenuTemplate) TableName() string {
	return "menu_templates"
}
