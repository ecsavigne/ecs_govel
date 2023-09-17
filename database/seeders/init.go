package seeders

import (
	"ecs_govel/config"
	"fmt"
)

type Seeders struct {
	*config.DbInstance
}

var Seeder = new(Seeders)

func init() {
	Seeder.DbInstance = &config.DbInstance{}
	if config.Database.DB != nil {
		Seeder.DbInstance = config.Database
	} else {
		fmt.Print("Erro nao é posivel criar o DBase object inicializador")
	}
}
