package models

import (
	"time"

	"gorm.io/gorm"
)

type Usuario struct {
	ID        int            `gorm:"primarykey;column:ci;type:varchar(11); not null"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at;type:datetime;default:null"`
	Nombre    string         `gorm:"column:nombre;type:vasrchar(35)"`
	Psassword string         `gorm:"column:password;type:vasrchar(35)"`
}
