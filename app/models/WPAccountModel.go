package models

import (
	"errors"
	"os"
	"strings"

	"new_whatsmeow/database/migrations"

	"github.com/joho/godotenv"
)

func (m *WPAccountModels) CreateNewWPAccount(host string, CompanyPhone string, app migrations.Application, webhook string) error {
	if m.WPNumberExists(CompanyPhone) {
		return errors.New("Exists")
	}

	wpaccount := migrations.WPAccounts{
		App:         app,
		WhatsappDns: host,
		PhoneNumber: CompanyPhone,
		WebhookText: webhook + "RPI/reciveTextMessage",
		WebhookFile: webhook + "RPI/reciveFileMessage",
		AppHash:     app.APIHash,
	}

	if !m.DB.NewRecord(wpaccount) { // Check if primary key is valid
		return m.DB.Error
	}

	m.DB.Create(&wpaccount)
	if m.DB.Error != nil {
		return m.DB.Error
	}

	return nil
}

func (m *WPAccountModels) WPNumberExists(CompanyPhone string) bool {
	var count int
	var wpaccount migrations.WPAccounts

	m.DB.Where("phone_number = ?", CompanyPhone).First(&wpaccount).Count(&count)
	return count == 1
}

func (m *WPAccountModels) GetWPAccount(CompanyPhone string) (migrations.WPAccounts, error) {
	var wpaccount migrations.WPAccounts

	m.DB.Where("phone_number = ?", CompanyPhone).First(&wpaccount)
	if m.DB.Error != nil {
		return wpaccount, m.DB.Error
	}

	return wpaccount, nil
}

func (m *WPAccountModels) GetWPAccounts(app migrations.Application) ([]migrations.WPAccounts, error) {
	var wpaccounts []migrations.WPAccounts

	m.DB.Model(&app).Related(&wpaccounts)
	if m.DB.Error != nil {
		return wpaccounts, m.DB.Error
	}

	return wpaccounts, nil
}

func (m *WPAccountModels) DeleteWPNumberSession(CompanyPhone string, app migrations.Application) error {
	if !m.WPNumberExists(CompanyPhone) {
		return errors.New("AccountNotRegistred")
	}
	_, err := ApplicationModel.GetApp(app.APIHash)
	if err != nil {
		return errors.New("AppNotRegistred")
	}

	wpaccount, err := m.GetWPAccount(CompanyPhone)
	if err != nil {
		return errors.New("ErrorRetrievingWpAccount")
	}
	m.DB.Unscoped().Delete(&wpaccount)
	return nil
}

func (m *WPAccountModels) GetAllWPAccounts() ([]migrations.WPAccounts, error) {
	godotenv.Load(".env")

	whatsappDns := "localhost:1337"
	if strings.Compare(os.Getenv("APP_ENV"), "production") == 0 {
		whatsappDns = "whatsmeowx" + m.AppID + ".socialhub.pro"
	}

	wpaccounts := make([]migrations.WPAccounts, 0)

	if err := m.DB.Where("whatsapp_dns = ?", whatsappDns).Order("created_at").Find(&wpaccounts).Error; err != nil {
		return wpaccounts, err
	}

	var accounts []migrations.WPAccounts

	accounts = append(accounts, wpaccounts...)
	wpaccounts = nil

	return accounts, nil
}
