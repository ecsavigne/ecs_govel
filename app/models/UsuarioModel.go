package models

import (
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nombre   string `gorm:"column:nombre;type:varchar(35)"`
	Password string `gorm:"column:password;type:varchar(35)"`
	Mail     string `gorm:"column:mail;type:varchar(100)"`
}
