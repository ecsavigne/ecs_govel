package models

import (
	"time"

	"gorm.io/gorm"
)

type Matriculas struct {
	ID           int            `gorm:"primarykey;column:id;autoIncrement; not null"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt    gorm.DeletedAt `gorm:"index;column:deleted_at;type:datetime;default:null"`
	CursoId      int            `gorm:"column:curso_id;default:null"`
	EstudianteCi string         `gorm:"column:estudiante_ci;type:varchar(11);default:null"`
	FechaIngreso time.Time      `gorm:"column:fecha_ingreso;type:datetime;default:null"`
}
