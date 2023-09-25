package models

import (
	"time"

	"gorm.io/gorm"
)

type Personas struct {
	CI              int              `gorm:"primarykey;column:ci; not null"`
	CreatedAt       time.Time        `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt       time.Time        `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt       gorm.DeletedAt   `gorm:"index;column:deleted_at;type:datetime;default:null"`
	Nombre          string           `gorm:"column:nombre;type:varchar(35);default:null"`
	Apellidos       string           `gorm:"column:apellidos;type:varchar(100);default:null"`
	Dir             string           `gorm:"column:dir;type:varchar(255);default:null"`
	Mail            string           `gorm:"column:mail;type:varchar(70);default:null"`
	Estudiante      Estudiantes      `gorm:"foreignKey:Ci"`
	InstructorCurso InstructorCursos `gorm:foreignKey:Ci`
}
