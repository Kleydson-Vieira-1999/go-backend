package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"           gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string    `json:"name"         gorm:"not null"`
	Email       string    `json:"email"        gorm:"unique;not null"`
	Picture     string    `json:"picture"`
	SSOProvider string    `json:"soo_provider" gorm:"type:sso_provider_enum;not null"`
	SSOID       string    `json:"sso_id"       gorm:"unique;not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
