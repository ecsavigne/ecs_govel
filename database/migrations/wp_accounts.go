package migrations

import (
	"github.com/jinzhu/gorm"
)

type WPAccounts struct {
	gorm.Model
	//WppID       uint
	App         Application `gorm:"foreignkey:AppRefer"` // use UserRefer as foreign key
	AppHash     string
	WhatsappDns string
	PhoneNumber string `gorm:"not null"`
	WebhookText string
	WebhookFile string
}
