package seeders

import (
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/database/migrations"
)

func (s *Seeders) CompaniesWhatsappSeeder() {
	err := s.DB.Exec("TRUNCATE TABLE company_whatsapps").Error
	if err != nil {
		logg.ErrorLogger.Println("\033[31m", err, "\033[0m")
	}

	companiesWhatsapps := migrations.CompanyWhatsapps{
		CompanyId:         1,
		CompanyWhatsappId: 1,
		Whatsapp:          "5521969537126",
	}
	s.DB.Create(&companiesWhatsapps)

	companiesWhatsapps = migrations.CompanyWhatsapps{
		CompanyId:         1,
		CompanyWhatsappId: 1,
		Whatsapp:          "552198261843",
	}
	s.DB.Create(&companiesWhatsapps)
}
