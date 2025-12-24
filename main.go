package main

import (
	//configs "ecs_govel/configs"

	// "ecs_govel/app/webhooks"
	config "ecs_govel/configs"
	// grpcserverinit "ecs_govel/grpcservice/app/server"
	_ "ecs_govel/routes"
)

func main() {
	//Test de Rutas
	// config.RouterList(config.GetEngine())
	//config.F_prepare_collection_postman("")

	config.AppRun()

	// Init GrpcService
	// grpcserverinit.InitGrpcService()
}
