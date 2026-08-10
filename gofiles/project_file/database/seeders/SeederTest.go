package seeders

import (
	"ecs_govel/database/migration"
	"fmt"
)

/*Protocolo para agregar un seeder*/
/*
func (s *Seeders) NameSeeder() {
	// Codigo del seeder
}
*/
// Ejemplo de agregar un seeder de prueba de aplicacion
func (s *Seeders) TestSeeder() {
	test := &migration.TestMigation{}

	if res := s.DB.FirstOrCreate(test); res.Error != nil {
		fmt.Println("Error creating test ", test, ". Error is: ", res.Error.Error())
		return
	}
}
