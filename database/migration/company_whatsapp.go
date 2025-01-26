package migration

import (
	"time"

	"gorm.io/gorm"
)

type CompanyWhatsapp struct {
	ID                   uint                 `json:"id" gorm:"primarykey"`
	CreatedAt            time.Time            `json:"created_at" gorm:"index"`
	UpdatedAt            time.Time            `json:"updated_at" gorm:"index"`
	DeletedAt            gorm.DeletedAt       `json:"deleted_at" gorm:"index"`
	CompanyId            uint64               `json:"company_id" gorm:"index"`
	CompanyWhatsappId    uint64               `json:"company_whatsapp_id" gorm:"index"`
	Whatsapp             string               `json:"whatsapp" gorm:"unique;not null"`
	ProcessGroupMessages bool                 `json:"process_group_messages"`
	CompanyWhatsappTags  []CompanyWhatsappTag `json:"company_whatsapp_tags" gorm:"foreignKey:CompanyWhatsappId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
