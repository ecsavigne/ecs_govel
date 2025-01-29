package model

import "gorm.io/gorm"

type ModelKernel interface {
	Insert(db *gorm.DB) *gorm.DB
	Update(db *gorm.DB) *gorm.DB
	Delete(db *gorm.DB) *gorm.DB
	// Get(db *gorm.DB, id uint)
}
