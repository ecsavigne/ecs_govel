package models

import (
	"time"

	"gorm.io/gorm"
)

type Contenido struct {
	ID                    int                   `gorm:"primaryKey;column:id;autoIncrement; not null"`
	CreatedAt             time.Time             `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt             time.Time             `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt             gorm.DeletedAt        `gorm:"index;column:deleted_at;type:datetime;default:null"`
	Tema                  string                `gorm:"column:tema;index;type:varchar(255);default:null"`
	CursosContenidos      []CursoContenido      `gorm:"foreignKey:ContenidoId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EstructurasContenidos []EstructuraContenido `gorm:"foreignKey:ContenidoId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
