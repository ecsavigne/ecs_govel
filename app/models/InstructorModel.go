package models

import (
	"time"

	"gorm.io/gorm"
)

type Instructore struct {
	CI                string            `gorm:"primaryKey;column:ci;type:varchar(11); not null"`
	CreatedAt         time.Time         `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt         time.Time         `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt         gorm.DeletedAt    `gorm:"index;column:deleted_at;type:datetime;default:null"`
	InstructorsCursos []InstructorCurso `gorm:"foreignKey:InstructorCi;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
