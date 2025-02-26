package seeders

import (
	model "ecs_govel/app/model"
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
	test := &model.TestModel{}

	if res := s.DB.FirstOrCreate(test); res.Error != nil {
		fmt.Println("Error creating test ", test, ". Error is: ", res.Error.Error())
		return
	}
}
