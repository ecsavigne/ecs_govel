package model

import (
	logecs "github.com/ecsavigne/logecs/log"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	Log logecs.Logger
)

func SetDB(db_ *gorm.DB) {
	db = db_
}

type ModelKernel interface {
	Insert() *gorm.DB
	Update() *gorm.DB
	Delete() *gorm.DB
	// Obtener por id
	GetById(id uint) ModelKernel
	// Obtener todos o limit
	Gets(...int) []ModelKernel
}

var SetGlobalDB = func(db_ *gorm.DB) {
	db = db_
}

var GetGlobalDB = func() *gorm.DB {
	return db
}
