package model

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type ModelKernel interface {
	Insert() *gorm.DB
	Update() *gorm.DB
	Delete() *gorm.DB
	// Obtener por id
	GetById(db *gorm.DB, id uint) ModelKernel
	// Obtener todos o limit
	Gets(...int) []ModelKernel
}

var SetGlobalDB = func(db_ *gorm.DB) {
	db = db_
}

var GetGlobalDB = func() *gorm.DB {
	return db
}
