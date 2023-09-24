package models

import (
	"time"

	"gorm.io/gorm"
)

type Contenidos struct {
	ID        int `gorm:"primarykey;column:id;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index;"`
	Tema      string         `gorm:"index;type:varchar(255);default:null"`
}
