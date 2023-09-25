package models

import (
	"time"

	"gorm.io/gorm"
)

type Estudiantes struct {
	CI          string         `gorm:"primarykey;column:ci;type:varchar(11);not null"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt   gorm.DeletedAt `gorm:"index;column:deleted_at;type:datetime;default:null"`
	si_juridico int            `gorm:"column:si_juridico;type:tinyint(1);default:null"`
	Matriculass []Matriculas   `gorm:foreignKey:EstudianteCi`
}
