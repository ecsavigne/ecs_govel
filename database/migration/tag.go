package migration

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID                  uint                 `json:"id" gorm:"primarykey"`
	CreatedAt           time.Time            `json:"created_at" gorm:"index"`
	UpdatedAt           time.Time            `json:"updated_at" gorm:"index"`
	DeletedAt           gorm.DeletedAt       `json:"deleted_at" gorm:"index"`
	ColorID             uint                 `json:"color_id" gorm:"index"`
	Color               Color                `json:"color"`
	Name                string               `json:"name"`
	ChatTags            []ChatTag            `json:"chat_tags" gorm:"foreignKey:TagID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CompanyWhatsappTags []CompanyWhatsappTag `json:"company_whatsapp_tags" gorm:"foreignKey:TagID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ContactTags         []ContactTag         `json:"contact_tags" gorm:"foreignKey:TagID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
