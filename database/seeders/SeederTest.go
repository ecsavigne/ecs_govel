package seeders

import (
	"fmt"
	"oficial_gin/app/models"
)

/*Protocolo para agregar un seeder*/
/*
func (s *Seeders) NameSeeder() {
	// Codigo del seeder
}
*/
// Ejemplo de agregar un seeder de prueba de aplicacion
func (s *Seeders) TestSeeder() {
	application := models.Application{
		AppName: "socialhub_chat",
		// AppHash:     "S4h_EPRZm-b46kyoUbUJ",
		WebhookText: "http://principal.socialhub.local/reciveTextMessage",
		WebhookFile: "http://principal.socialhub.local/reciveFileMessage",
	}

	//TODO db.DB.Create(&application) intentaba criar y si ya existia el registro devolvia error de primary key  ERROR: duplicate key value violates unique constraint "applications_api_hash_key"
	if res := s.DB.FirstOrCreate(&application); res.Error != nil {
		fmt.Println("Error creating application ", application, ". Error is: ", res.Error.Error())
		return
	}
}
