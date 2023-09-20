package migrations

import (
	"github.com/jinzhu/gorm"
)

// NOTA: Las migration tendaran el Nombre <Mxxxxxxx>
type Test struct {
	gorm.Model
	// Owner       User
	// AppName     string `gorm:"not null"`
	// APIHash     string `gorm:"unique"`
	// WebhookText string
	// WebhookFile string
}
