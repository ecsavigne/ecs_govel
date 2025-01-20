package migration

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"index"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ContactTags []ContactTag   `json:"contact_tags" gorm:"foreignKey:ContactID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
