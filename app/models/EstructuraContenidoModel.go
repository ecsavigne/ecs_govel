package models

import (
	"time"

	"gorm.io/gorm"
)

type EstructuraContenido struct {
	ID                       int            `gorm:"primaryKeys;column:id;autoIncrement; not null"`
	CreatedAt                time.Time      `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt                time.Time      `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt                gorm.DeletedAt `gorm:"index;column:deleted_at;type:datetime;default:null"`
	CantidadClasePlane       int            `gorm:"column:cantidad_clase_plane;default:null"`
	CantidadTematicaPrograma int            `gorm:"column:cantidad_tematica_programa;default:null"`
	ContenidoId              int            `gorm:"column:contenido_id"`
	EstructuraId             int            `gorm:"column:estructura_id"`
}
