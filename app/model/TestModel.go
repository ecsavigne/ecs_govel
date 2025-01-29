package model

import (
	"ecs_govel/database/migration"

	"gorm.io/gorm"
)

type Test migration.Test

/*Insert(db *gorm.DB) *gorm.DB
Update(db *gorm.DB) *gorm.DB
Delete(db *gorm.DB) *gorm.DB*/

func (*Test) TableName() string {
	return "Tests"
}

func (*Test) GetModelName() string {
	return "Test"
}

func (*Test) Insert() *gorm.DB {
	return db.Create(&Test{})
}

func (*Test) Update() *gorm.DB {
	return db.Save(&Test{})
}

func (*Test) Delete() *gorm.DB {
	return db.Delete(&Test{})
}
