package models

import (
	"time"

	"gorm.io/gorm"
)

type Cursos struct {
	ID                int                `gorm:"primarykey;column:id;autoIncrement; not null"`
	CreatedAt         time.Time          `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt         time.Time          `gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt         gorm.DeletedAt     `gorm:"index;column:deleted_at;type:datetime;default:null"`
	FechaIngreso      time.Time          `gorm:"type:datetime;column:fecha_ingreso"`
	DuracionHora      time.Time          `gorm:"type:datetime;column:duracion_hora"`
	SiCertificado     int                `gorm:"column:si_certificado; type:tinyint(1); default:null"`
	InstructorsCursos []InstructorCursos `gorm:foreignKey:CursoId`
	CursosContenidos  []CursoContenidos  `gorm:foreignKey:CursoId`
	Matriculass       []Matriculas       `gorm:foreignKey:CursoId`
}
