package models

import (
	"time"

	"gorm.io/gorm"
)

type Curso struct {
	ID                int               `gorm:"primaryKey;column:id;autoIncrement; not null"`
	CreatedAt         time.Time         `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt         time.Time         `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt         gorm.DeletedAt    `gorm:"index;column:deleted_at;type:datetime;default:null"`
	FechaCreacion     time.Time         `gorm:"type:datetime;column:fecha_creacion"`
	DuracionHora      int               `gorm:"type:int;column:duracion_hora"`
	SiCertificado     bool              `gorm:"column:si_certificado;type:bool;default:false"`
	InstructorsCursos []InstructorCurso `gorm:"foreignKey:CursoId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CursosContenidos  []CursoContenido  `gorm:"foreignKey:CursoId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Matriculass       []Matricula       `gorm:"foreignKey:CursoId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
