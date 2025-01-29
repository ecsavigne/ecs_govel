package model

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type ModelKernel interface {
	Insert() *gorm.DB
	Update() *gorm.DB
	// Delete(db *gorm.DB) *gorm.DB
	Delete() *gorm.DB
	// Get(db *gorm.DB, id uint)
}

var SetGlobalDB = func(db_ *gorm.DB) {
	db = db_
}

var GetGlobalDB = func() *gorm.DB {
	return db
}
