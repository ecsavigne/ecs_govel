package main

import (
	//configs "ecs_govel/configs"

	// "ecs_govel/app/webhooks"
	config "ecs_govel/configs"
	_ "ecs_govel/routes"
)

func main() {
	//Test de Rutas
	// config.RouterList(config.GetEngine())
	//config.F_prepare_collection_postman("")
	config.AppRun()
}
