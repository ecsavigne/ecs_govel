package models

import "ecs_govel/config"

// Definicion de la estructura de modelo que carga todos los metodos del
// ORM GORM en db
type NombreModels struct {
	*config.DbInstance
}

//var NombreModel = new(NombreModels)

// func init() {
// 	NombreModel.DbInstance = config.Database
// }
