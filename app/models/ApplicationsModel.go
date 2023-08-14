package models

import (
	"new_whatsmeow/database/migrations"
)

func (m *ApplicationModels) GetApps() ([]migrations.Application, error) {
	// GetApp : Find app by apihash
	var applications []migrations.Application

	m.DB.Find(&applications)
	if m.DB.Error != nil {
		return applications, m.DB.Error
	}
	return applications, nil
}

// GetApp : Find app by apihash
func (m *ApplicationModels) GetApp(apihash string) (migrations.Application, error) {

	var application migrations.Application

	m.DB.Where("api_hash = ?", apihash).First(&application)
	if m.DB.Error != nil {
		return application, m.DB.Error
	}
	return application, nil
}

// AppIsAuth : Check if app hash is valid
func (m *ApplicationModels) AppIsAuth(apihash string) bool {

	var application migrations.Application
	var count int

	m.DB.Where("api_hash = ?", apihash).First(&application).Count(&count)
	return count == 1
}
