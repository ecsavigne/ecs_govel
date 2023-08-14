package migrations

import (
	"github.com/jinzhu/gorm"
)

type Sessions struct {
	gorm.Model
	PhoneNumber string `gorm:"not null"`
	ClientID    string
	ClientToken string
	ServerToken string
	EncKey      []byte
	MacKey      []byte
	Wid         string
}
