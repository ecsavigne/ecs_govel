package models

import (
	"time"

	"gorm.io/gorm"
)

type InstructorCurso struct {
	ID           int            `gorm:"primaryKey;column:id;autoIncrement; not null"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt    gorm.DeletedAt `gorm:"index;column:deleted_at;type:datetime;default:null"`
	InstructorCi string         `gorm:"column:instructor_ci;type:varchar(11);default:null;"`
	CursoId      int            `gorm:"column:curso_id;default:null"`
}
