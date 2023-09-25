package models

import (
	"time"

	"gorm.io/gorm"
)

type InstructorCursos struct {
	ID           int            `gorm:"primarykey;column:id;autoIncrement; not null"`
	CreatedAt    time.Time      `gorm:"column:s;type:datetime;default:null"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt    gorm.DeletedAt `gorm:"index;column:deleted_at;type:datetime;default:null"`
	InstructorCi string         `gorm:"column:instructor_ci;type:varchar(11);default:null;"`
	CursoId      int            `gorm:"column:curso_id;default:null"`
}
