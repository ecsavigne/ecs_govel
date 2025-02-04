package model

import (
	"ecs_govel/database/migration"

	"gorm.io/gorm"
)

type TestModel migration.TestMigation

/*
Insert() *gorm.DB
	Update() *gorm.DB
	Delete() *gorm.DB
	// Obtener por id
	GetById(db *gorm.DB, id uint) ModelKernel
	// Obtener todos o limit
	Gets(...int) []ModelKernel
*/

func (*TestModel) TableName() string {
	return "TestModels"
}

func (*TestModel) GetModelName() string {
	return "TestModel"
}

func (*TestModel) Insert() *gorm.DB {
	return db.Create(&TestModel{})
}

func (*TestModel) Update() *gorm.DB {
	return db.Save(&TestModel{})
}

func (*TestModel) Delete() *gorm.DB {
	return db.Delete(&TestModel{})
}

func (t *TestModel) GetById(db *gorm.DB, id uint) ModelKernel {
	return nil
}

func (t *TestModel) Gets(limit ...int) []ModelKernel {
	return nil
}
