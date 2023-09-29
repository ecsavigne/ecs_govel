package models

import (
	"time"

	"gorm.io/gorm"
)

type Persona struct {
	CI         string         `gorm:"primaryKey;column:ci;type:varchar(11); not null"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt  gorm.DeletedAt `gorm:"index;column:deleted_at;type:datetime;default:null"`
	Nombre     string         `gorm:"column:nombre;type:varchar(35);default:null"`
	Apellidos  string         `gorm:"column:apellidos;type:varchar(100);default:null"`
	Dir        string         `gorm:"column:dir;type:varchar(255);default:null"`
	Mail       string         `gorm:"column:mail;type:varchar(70);default:null"`
	Estudiante Estudiante     `gorm:"foreignKey:CI;references:CI;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Instructor Instructore    `gorm:"foreignKey:CI;references:CI;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
