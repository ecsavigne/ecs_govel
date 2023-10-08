package seeders

import (
	"ecs_govel/config/db"
	"fmt"
)

type Seeders struct {
	*db.DbInstance
}

var Seeder = new(Seeders)

func init() {
	Seeder.DbInstance = &db.DbInstance{}
	if db.Orm.DB != nil {
		Seeder.DbInstance = db.Orm
	} else {
		fmt.Print("Erro nao é posivel criar o DBase object inicializador")
	}
}
