package migration

import (
	"time"

	"gorm.io/gorm"
)

// // User  : User model
type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Username  string         `gorm:"not null;unique"`
	APIHash   string         `gorm:"unique"`
}
