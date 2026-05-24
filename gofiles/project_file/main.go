package main

import (
	//configs "ecs_govel/configs"

	// "ecs_govel/rest/app/webhooks"
	"ecs_govel/command"
	config "ecs_govel/configs"
	grpcserverinit "ecs_govel/grpcservice/app/server"
	_ "ecs_govel/routes"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func run() {
	// commands
	go command.CreateJobNameXserivePrometheus()

	//Test de Rutas
	// config.RouterList(config.GetEngine())
	//config.F_prepare_collection_postman("")

	go config.AppRun()
	time.Sleep(time.Duration(2) * time.Second)

	go command.RunJobs()

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
