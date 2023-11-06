package seeders

import (
	"reflect"

	"gorm.io/gorm"
)

type Seeders struct {
	*gorm.DB
	// arreglo con el nombre de todos los seeder a ejecutar seeder = SeederNombre, file SeederNombre.go
	//function =  (s * Seeders) SeederNombre() {}
	SeedersLoad []string
}

var Seeder = new(Seeders)

// Solo para inicializar el arreglo con el nombre de los seeder a ejecutar el que no este aqui no ejecuta
func init() {
	Seeder.SeedersLoad = append(Seeder.SeedersLoad, "SeederUsuario")
}

func Init(c *gorm.DB) {
	Seeder.DB = c
}

// Ejecuta un metodo del objeto seeder pasando como parametro el seeder y el nombre del metodo
// retorna true si ejecuta false si no seencuentra
func EjecutarMetodo(seederObj interface{}, nombreMetodoSeeder string) bool {
	seederInstance := reflect.ValueOf(seederObj)
	metodoSeeder := seederInstance.MethodByName(nombreMetodoSeeder)

	if metodoSeeder.IsValid() {
		metodoSeeder.Call(nil)
		return true
	} else {
		return false
	}
}
