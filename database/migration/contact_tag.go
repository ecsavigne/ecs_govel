package migration

import (
	"time"

	"gorm.io/gorm"
)

type ContactTag struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	TagID     uint           `json:"tag_id" gorm:"index"`
	Tag       Tag            `json:"tag"`
	ContactID uint           `json:"contact_id" gorm:"index"`
	Contact   Contact        `json:"contact"`
}
