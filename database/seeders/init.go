package seeders

import (
	"gorm.io/gorm"
)

type Seeders struct {
	*gorm.DB
}

var Seeder = new(Seeders)

func Init(c *gorm.DB) {
	// Seeder.DbInstance = &db.DbInstance{}
	// if db.Orm.DB != nil {
	// 	Seeder.DbInstance = db.Orm
	// } else {
	// 	fmt.Print("Erro nao é posivel criar o DBase object inicializador")
	// }
	Seeder.DB = c
}
