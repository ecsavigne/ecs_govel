package models

import (
	"time"

	"gorm.io/gorm"
)

type Estructuras struct {
	ID                    int                    `gorm:"primarykey;column:id;autoIncrement; not null"`
	CreatedAt             time.Time              `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt             time.Time              `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt             gorm.DeletedAt         `gorm:"index;column:deleted_at;type:datetime;default:null"`
	TipoEstructura        string                 `gorm:"column:tipo_estructura;type:varchar(70);default:null"`
	EstructurasContenidos []EstructuraContenidos `gorm:EstructuraId`
}
