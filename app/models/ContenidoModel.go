package models

import (
	"time"

	"gorm.io/gorm"
)

type Contenidos struct {
	ID                    int                    `gorm:"primarykey;column:id;autoIncrement; not null"`
	CreatedAt             time.Time              `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt             time.Time              `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt             gorm.DeletedAt         `gorm:"index;column:deleted_at;type:datetime;default:null"`
	Tema                  string                 `gorm:"column:tema;index;type:varchar(255);default:null"`
	CursosContenidos      []CursoContenidos      `gorm:foreignKey:ContenidoId`
	EstructurasContenidos []EstructuraContenidos `gorm:foreignKey:ContenidoId`
}
