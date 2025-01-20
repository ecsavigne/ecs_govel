package migration

import "gorm.io/gorm"

// Sessions :
type Session struct {
	gorm.Model
	PhoneNumber string `gorm:"not null"`
	ClientID    string
	ClientToken string
	ServerToken string
	EncKey      []byte
	MacKey      []byte
	Wid         string
}
