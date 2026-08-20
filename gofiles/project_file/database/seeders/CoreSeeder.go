package seeders

import (
	"reflect"

	"gorm.io/gorm"
)

type seeders struct {
	*gorm.DB
}

var seeder = &seeders{}

func ExecuteSeeders(dBase *gorm.DB) {
	seeder.DB = &gorm.DB{}

	// Executa seeder
	typ := reflect.TypeFor[*seeders]() // information of struct
	value := reflect.ValueOf(&seeder)  // value of struct

	for method := range typ.Methods() {
		// if method.Name == "TestSeeder" {
		// 	continue
		// }
		value.MethodByName(method.Name).Call([]reflect.Value{})
	}
}
