package seeder

import (
	"ecs_govel/pkg/pkglog"
	"reflect"

	"gorm.io/gorm"
)

type seeders struct {
	DB *gorm.DB
}

var seeder = &seeders{}

func ExecuteSeeders(dBase *gorm.DB) {
	seeder.DB = dBase

	// Executa seeder
	typ := reflect.TypeFor[*seeders]() // information of struct
	value := reflect.ValueOf(seeder)   // value of struct

	for method := range typ.Methods() {
		pkglog.Log.Infof("executing seeders Method: %+v\n", method)

		methodValue := value.MethodByName(method.Name)
		if method.Func.IsValid() {
			methodValue.Call([]reflect.Value{})
		}

		if method.Name == "TestSeeder" {
			continue
		}
	}
}
