package migrations

import (
	"github.com/jinzhu/gorm"
)

type History struct {
	gorm.Model
	MessageId      string `gorm:"unique"`
	PhoneNumber    string
	Messages       string
	WhatsappDate   string
	ServerResponse string
	MessageType    string
	HandleErr      string
	ServiceDate    string
	JsonPayload    string
}
