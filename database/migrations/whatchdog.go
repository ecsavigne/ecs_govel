package migrations

import (
	"github.com/jinzhu/gorm"
)

type WhatchDog struct {
	gorm.Model
	MessageId   string
	PhoneNumber string
	RemoteJid   string
	Log         string
	Error       bool
}
