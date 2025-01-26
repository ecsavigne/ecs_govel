package migration

import (
	"time"

	"gorm.io/gorm"
)

// WPAccounts : WhatsAppAccount model (test if FK is working otherwise do it manually)
type WPAccount struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	//WppID       uint
	//TODONo hace falta solo proboca un error por eso comente App. Tuve que comentar en database.go->CreateNewWPAccount en la linea 217 atributo App
	//App         Application `gorm:"foreignkey:AppRefer"` // use UserRefer as foreign key
	AppHash     string `json:"app_hash" gorm:"not null"`
	WhatsappDns string `json:"whatsapp_dns" gorm:"not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null; index"`
	WebhookText string `json:"webhook_text"`
	WebhookFile string `json:"webhook_file"`
}
