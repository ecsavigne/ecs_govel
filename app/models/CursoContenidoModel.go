package models

import (
	"time"

	"gorm.io/gorm"
)

type CusrsoContenidos struct {
	ID        int `gorm:"primarykey;column:id;autoIncrement; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index;"`
	Tema      string         `gorm:"index;type:varchar(255);default:null"`
}
