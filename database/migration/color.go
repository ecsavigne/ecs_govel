package migration

import (
	"time"

	"gorm.io/gorm"
)

type Color struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"index"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ColorNumber int            `json:"color_number"`
	Rgb         string         `json:"rgb"`
	HexaDecimal string         `json:"hexa_decimal"`
	Tags        []Tag          `json:"tags" gorm:"foreignKey:ColorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
