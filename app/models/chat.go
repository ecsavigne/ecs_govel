//lint:file-ignore ST1005 Ignore capitalized strings error

package models

import (
	"fmt"
	"oficial_gin/database/migration"

	"gorm.io/gorm"
)

// Chat :
type Chat migration.Chat

func (c *Chat) Create(db *gorm.DB) error {

	if res := db.Create(&c); res.Error != nil {
		return fmt.Errorf("Error in model.Chat.Create:function. Chat values: %+v. Error was: %s", c, db.Error.Error())
		// return db.DB.Error
	}
	return db.Error
}

func (c *Chat) UpdateCompanyPhone(db *gorm.DB, CompanyId, CompanyWhatsappId uint64) *gorm.DB {
	return db.Model(c).
		Where("company_id = ? AND company_whatsapp_id = ?", CompanyId, CompanyWhatsappId).
		Update("company_phone", c.CompanyPhone)
}

// Copy chat for companyWhatsapp
// params {
//
//			"CompanyId" 			: int64,
//			"CompanyWhatsappIdOld"	: int64,
//			"CompanyWhatsappIdNew"	: int64,
//			"CompanyWhatsappOld"	: string,
//			"CompanyWhatsappNew"	: string
//	}
func (c *Chat) CopyChatForCompanyWhatsapp(db *gorm.DB, params map[string]interface{}) *gorm.DB {
	companyId := params["CompanyId"].(int64)
	companyWhatsappIdOld := params["CompanyWhatsappIdOld"].(int64)
	companyWhatsappIdNew := params["CompanyWhatsappIdNew"].(int64)
	companyWhatsappOld := params["CompanyWhatsappOld"].(string)
	companyWhatsappNew := params["CompanyWhatsappNew"].(string)
	arrayContacts := params["ArrayContacts"].([]string)
	condition := "company_id = ? and company_whatsapp_id = ? and company_phone = ? "
	if len(arrayContacts) > 0 {
		condition += "and contact_phone in ("
		i := 0
		for _, v := range arrayContacts {
			if i != 0 {
				condition += ","
			}
			condition += "'" + v + "'"
			i++
		}
		condition += ")"
	}

	var chat []Chat
	res := db.Table("chat").
		Where(condition, companyId, companyWhatsappIdOld, companyWhatsappOld).
		FindInBatches(&chat, 100, func(tx *gorm.DB, batch int) error {
			for _, chat := range chat {
				// Crear Chat new
				chatTemp := new(Chat)
				*chatTemp = chat
				chatTemp.ID = 0
				chatTemp.CompanyPhone = companyWhatsappNew
				chatTemp.CompanyWhatsappId = uint64(companyWhatsappIdNew)
				chatTemp.Create(tx)
			}
			return nil
		})

	if err := res.Error; err != nil {
		fmt.Printf("Erro ao processar aos chat: %s\n", err.Error())
		return nil
	}
	return res
}
