package migrations

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserID   uint
	Username string `gorm:"not null;unique"`
	APIHash  string `gorm:"unique"`
}
