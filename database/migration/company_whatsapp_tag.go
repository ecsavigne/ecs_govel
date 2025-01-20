package migration

import (
	"time"

	"gorm.io/gorm"
)

type CompanyWhatsappTag struct {
	ID                uint            `json:"id" gorm:"primarykey"`
	CreatedAt         time.Time       `json:"created_at" gorm:"index"`
	UpdatedAt         time.Time       `json:"updated_at" gorm:"index"`
	DeletedAt         gorm.DeletedAt  `json:"deleted_at" gorm:"index"`
	CompanyWhatsappId uint            `json:"company_whatsapp_id" gorm:"index"`
	CompanyWhatsapp   CompanyWhatsapp `json:"company_whatsapp"`
	TagID             uint            `json:"tag_id" gorm:"index"`
	Tag               Tag             `json:"tag"`
}
