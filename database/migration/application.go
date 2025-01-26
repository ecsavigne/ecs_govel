package migration

import (
	"time"

	"gorm.io/gorm"
)

type Application struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"index"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	AppName     string         `gorm:"not null" json:"appname"`
	APIHash     string         `gorm:"unique" json:"apihash"`
	WebhookText string         `json:"webhooktext"`
	WebhookFile string         `json:"webhookfile"`
	WPAccounts  []WPAccount    `json:"wp_accounts"  gorm:"foreignKey:AppHash;references:APIHash;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
