package migrations

import (
	"github.com/jinzhu/gorm"
)

type Application struct {
	gorm.Model
	Owner       User
	AppName     string `gorm:"not null"`
	APIHash     string `gorm:"unique"`
	WebhookText string
	WebhookFile string
}
