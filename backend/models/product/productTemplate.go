package product

import (
	"github.com/google/uuid"
)

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
