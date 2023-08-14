package models

import (
	"new_whatsmeow/database/migrations"
)

func (m *CompanyWhatsappModels) CreateCompanyWhatsapp(companyWhatsapp migrations.CompanyWhatsapps) error {
	if !m.DB.NewRecord(companyWhatsapp) {
		return m.DB.Error
	}

	var found migrations.CompanyWhatsapps
	m.DB.Where("company_whatsapp_id = ? and whatsapp = ?", companyWhatsapp.CompanyWhatsappId, companyWhatsapp.Whatsapp).First(&found)
	if found.CompanyId != 0 {
		return nil
	}

	m.DB.Create(&companyWhatsapp)
	if m.DB.Error != nil {
		return m.DB.Error
	}

	return nil
}

func (m *CompanyWhatsappModels) GetCompanyWhatsapp(whatsapp string) (migrations.CompanyWhatsapps, error) {
	var companyWhatsapp migrations.CompanyWhatsapps
	m.DB.Where("whatsapp = ?", whatsapp).First(&companyWhatsapp)
	if m.DB.Error != nil {
		return companyWhatsapp, m.DB.Error
	}
	return companyWhatsapp, nil
}

func (m *CompanyWhatsappModels) UpdateCompanyWhatsapp(companyWhatsapp migrations.CompanyWhatsapps) error {
	if companyWhatsapp.Whatsapp != "" {
		m.DB.Model(&companyWhatsapp).Where("company_whatsapp_id = ?", companyWhatsapp.CompanyWhatsappId).Update("whatsapp", companyWhatsapp.Whatsapp)
		if m.DB.Error != nil {
			return m.DB.Error
		}
	}
	return nil
}

func (m *CompanyWhatsappModels) DeleteCompanyWhatsapp(companyWhatsapp migrations.CompanyWhatsapps) error {
	m.DB.Unscoped().Where("company_whatsapp_id = ?", companyWhatsapp.CompanyWhatsappId).Delete(&companyWhatsapp)
	if m.DB.Error != nil {
		return m.DB.Error
	}
	return nil
}

func (m *CompanyWhatsappModels) UpdateWPAccount(wpAccounts migrations.WPAccounts, whatsapp string) error {
	m.DB.Model(&wpAccounts).Where("company_phone = ?", whatsapp).Update(&wpAccounts)
	if m.DB.Error != nil {
		return m.DB.Error
	}
	return nil
}
