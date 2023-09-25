package models

import (
	"time"

	"gorm.io/gorm"
)

type CursoContenidos struct {
	ID          int            `gorm:"primarykey;column:id;autoIncrement; not null"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt   gorm.DeletedAt `gorm:"index;column:deleted_at;type:datetime;default:null"`
	ContenidoCi int            `gorm:"column:contenido_ci;default:null"`
	CursoId     int            `gorm:"column:curso_id;default:null"`
}
