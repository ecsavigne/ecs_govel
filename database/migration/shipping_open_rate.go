package migration

import (
	"time"

	"gorm.io/gorm"
)

type ShippingOpenRate struct {
	ID         uint           `json:"id" gorm:"primarykey"`
	CreatedAt  time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"index"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	MessageId  string         `json:"message_id" gorm:"unique;not null"`
	CompaingId uint64         `json:"compaing_id" gorm:"index;not null"`
	Readed     bool           `json:"readed"`
}
