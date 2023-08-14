package seeders

import (
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/database/migrations"
)

func (s *Seeders) ApplicationSeeder() {
	err := s.DB.Exec("TRUNCATE TABLE applications").Error
	if err != nil {
		logg.ErrorLogger.Println("\033[31m", err, "\033[0m")
	}
	application := &migrations.Application{
		AppName:     "socialhub_chat",
		APIHash:     "S4h_EPRZm-b46kyoUbUJ",
		WebhookText: "http://principal.socialhub.local/reciveTextMessage",
		WebhookFile: "http://principal.socialhub.local/reciveFileMessage",
	}
	s.DB.Create(&application)
}
