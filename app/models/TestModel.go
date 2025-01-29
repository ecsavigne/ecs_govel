package model

import (
	"oficial_gin/database/migration"

	"gorm.io/gorm"
)

type Audio migration.Audio

/*Insert(db *gorm.DB) *gorm.DB
Update(db *gorm.DB) *gorm.DB
Delete(db *gorm.DB) *gorm.DB*/

func (*Audio) TableName() string {
	return "audios"
}

func (*Audio) GetModelName() string {
	return "Audio"
}

func (*Audio) Insert(db *gorm.DB) *gorm.DB {
	return db.Create(&Audio{})
}

func (*Audio) Update(db *gorm.DB) *gorm.DB {
	return db.Save(&Audio{})
}

func (*Audio) Delete(db *gorm.DB) *gorm.DB {
	return db.Delete(&Audio{})
}
