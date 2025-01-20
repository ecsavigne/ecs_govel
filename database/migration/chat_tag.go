package migration

import (
	"time"

	"gorm.io/gorm"
)

type ChatTag struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	TagID     uint           `json:"tag_id"`
	ChatID    uint           `json:"chat_id"`
	Tag       Tag            `json:"tag"`
	Chat      Chat           `json:"chat"`
}
