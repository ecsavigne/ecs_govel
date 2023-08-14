package migrations

import (
	"github.com/jinzhu/gorm"
)

type CompanyWhatsapps struct {
	gorm.Model
	CompanyId         uint64
	CompanyWhatsappId uint64
	Whatsapp          string
	// Whatsapp          string `gorm:"unique;not null"`
}
