package seeder

import (
	"ecs_govel/database/migration"
	"ecs_govel/pkg/pkglog"
	"fmt"
)

/*Protocolo para agregar un seeder*/
/*
func (s *Seeders) NameSeeder() {
	// Codigo del seeder
}
*/
// Ejemplo de agregar un seeder de prueba de aplicacion
func (s *seeders) TestSeeder() {
	test := &migration.TestMigation{}
	_ = test

	if res := s.DB.FirstOrCreate(test); res.Error != nil {
		fmt.Println("Error creating test ", test, ". Error is: ", res.Error.Error())
		return
	}
	pkglog.Log.Warnf("executed TestSeeder: %+v\n", s.DB)
}
