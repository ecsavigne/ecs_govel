package database

import (
	c_ "ecs_govel/configs"
	"ecs_govel/database/driver/drivergorm"
	"ecs_govel/database/driver/drivermongo"
	"ecs_govel/database/shared"
)

var (
	managerGormDB *shared.DBManager = &shared.DBManager{}

	GetGormDB = func() *shared.DBManager {
		return managerGormDB
	}
)

func init() {
	switch c_.DB_TYPE {
	case "postgres":
		drivergorm.PostgresDB(managerGormDB)
	case "mongo":
		drivermongo.MongoDB()
	case "all":
		drivermongo.MongoDB()
		drivergorm.PostgresDB(managerGormDB)
	default:
		drivergorm.PostgresDB(managerGormDB)
	}
}
