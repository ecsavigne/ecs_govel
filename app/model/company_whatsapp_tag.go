package model

import (
	"ecs_govel/database/migration"
	"fmt"

	"gorm.io/gorm"
)

type CompanyWhatsappTag migration.CompanyWhatsappTag

// Arg0: CompanyWhatsappId, Arg1: TagID,  Arg2: id of CompanyWhatsappTagID puede o no pasarse si no se pasa gorm da un id...,
func NewCompanyWhatsappTag(arg ...interface{}) (*CompanyWhatsappTag, error) {
	if len(arg) > 0 {
		if arg[0] == nil || arg[1] == nil {
			return nil, fmt.Errorf("Arg0 TagID, Arg1 ChatsID are required")
		}

		CompanyWhatsappId, ok := arg[0].(uint)
		if !ok {
			return nil, fmt.Errorf("Arg0 CompanyWhatsappId is not uint")
		}

		TagID, ok := arg[1].(uint)
		if !ok {
			return nil, fmt.Errorf("Arg1 TagID is not uint")
		}

		if arg[2] != nil {
			id, ok := arg[2].(uint)
			if !ok {
				return nil, fmt.Errorf("Arg2 CompanyWhatsappTagID is not uint")
			}
			return &CompanyWhatsappTag{
				CompanyWhatsappId: CompanyWhatsappId,
				TagID:             TagID,
				ID:                id,
			}, nil
		} else {
			return &CompanyWhatsappTag{
				CompanyWhatsappId: CompanyWhatsappId,
				TagID:             TagID,
			}, nil
		}
	} else {
		return &CompanyWhatsappTag{}, nil
	}
}

func (cwt *CompanyWhatsappTag) Insert(db *gorm.DB) *gorm.DB {
	return db.Create(cwt)
}

func (cwt *CompanyWhatsappTag) Update(db *gorm.DB) *gorm.DB {
	return db.Save(cwt)
}

// Param2: CompanyWhatsappId, Param3: TagID
func (cwt *CompanyWhatsappTag) Get(db *gorm.DB, CompanyWhatsappId, TagID uint) {
	db.First(cwt, "refer_company_whatsapp_id = ? AND tag_id = ?", CompanyWhatsappId, TagID)
}
