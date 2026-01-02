package main

import (
	//configs "ecs_govel/configs"

	// "ecs_govel/app/webhooks"
	// "ecs_govel/app/console"
	config "ecs_govel/configs"
	grpcserverinit "ecs_govel/grpcservice/app/server"
	_ "ecs_govel/routes"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func run() {
	// commands
	// go console.CreateJobNameXserivePrometheus()

	//Test de Rutas
	// config.RouterList(config.GetEngine())
	//config.F_prepare_collection_postman("")

	go config.AppRun()

	// Init GrpcService
	if strings.ToLower(config.TYPE_SERVICES) == "grpc" {
		grpcserverinit.InitGrpcService()
	} else {
		wait := make(chan os.Signal, 1)
		signal.Notify(wait, os.Interrupt, syscall.SIGTERM)
		<-wait
	}
}

func main() {
	run()
}
