package main

import (
	//configs "ecs_govel/configs"

	"fmt"
	// "ecs_govel/app/webhooks"
	config "ecs_govel/configs"
	_ "ecs_govel/routes"
)

func main() {
	// fmt.Println("Host:", configs.HTTP_SERVER_HOST)
	//Test de Rutas
	// config.RouterList(config.GetEngine())
	//config.F_prepare_collection_postman("")
	config.AppRun()
	fmt.Println("Starting Webhooks")
	//new(webhooks.WPSession).RestQueue()
	fmt.Println("Starting Webhooks 2")
}
