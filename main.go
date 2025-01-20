package main

import (
	//configs "oficial_gin/configs"

	"fmt"
	// "oficial_gin/app/webhooks"
	config "oficial_gin/configs"
	_ "oficial_gin/routes"
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
