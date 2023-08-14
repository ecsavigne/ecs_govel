package models

import (
	"new_whatsmeow/database/migrations"
)

func (m *WhatchDogModels) CreteWhatchDog(Function string, CompanyPhone string, ContactPhone string, Log string, Error bool) error {
	message := migrations.WhatchDog{
		PhoneNumber: CompanyPhone,
		RemoteJid:   ContactPhone,
		Log:         Log,
		Error:       Error,
	}
	if !m.DB.NewRecord(message) {
		return m.DB.Error
	}
	m.DB.Create(&message)
	if m.DB.Error != nil {
		return m.DB.Error
	}
	return nil
}

func (m *WhatchDogModels) SaveMessageIntoWhatchDog(messageid string, CompanyPhone string, ContactPhone string, log string, error bool) error {
	message := migrations.WhatchDog{
		PhoneNumber: CompanyPhone,
		RemoteJid:   ContactPhone,
		Log:         log,
		Error:       error,
	}
	if !m.DB.NewRecord(message) {
		return m.DB.Error
	}

	m.DB.Create(&message)
	if m.DB.Error != nil {
		return m.DB.Error
	}

	return nil
}

func (m *WhatchDogModels) LoadFromWhatchDog(CompanyPhone string, dateInit string, dateEnd string) []migrations.WhatchDog {
	var whatchDog []migrations.WhatchDog
	m.DB.Raw("SELECT * from whatch_dogs where phone_number like ? and created_at >= ? and created_at <= ?  order by created_at asc", CompanyPhone, dateInit, dateEnd).Scan(&whatchDog)
	return whatchDog
}
