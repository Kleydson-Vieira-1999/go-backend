package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"           gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID      uuid.UUID `json:"user_id"      gorm:"type:uuid;not null"`
	Name        string    `json:"name"         gorm:"not null"`
	Description string    `json:"description"  `
	CostPrice   int       `json:"cost_price"   gorm:"not null"`
	Price       int       `json:"price"        gorm:"not null"`
	ImageBase64 string    `json:"image"        gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Product) TableName() string {
	return "products"
}

type ProductWithAvailability struct {
	Product
	IsAvailable bool `json:"is_available" gorm:"column:is_available"`
}

type ProductTemplate struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name         string    `gorm:"not null"`
	Description  string
	CostPrice    int       `gorm:"not null"`
	Price        int       `gorm:"not null"`
	ImageBase64  string    `gorm:"type:text"`
	UpdatedAt    string
}

func (ProductTemplate) TableName() string {
	return "product_templates"
}
