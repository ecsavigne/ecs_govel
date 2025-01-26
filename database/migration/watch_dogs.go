package migration

import (
	"time"

	"gorm.io/gorm"
)

// WhatchDog :
type WhatchDog struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"index"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	MessageId   string         `gorm:"unique" json:"message_id"`
	PhoneNumber string         `json:"phone_number"`
	RemoteJid   string         `json:"remote_jid"`
	Log         string         `json:"log"`
	Error       bool           `json:"error"`
}
