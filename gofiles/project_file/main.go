package main

import (
	//configs "ecs_govel/configs"

	// "ecs_govel/rest/app/webhooks"
	"ecs_govel/command"
	config "ecs_govel/configs"
	_ "ecs_govel/database"
	grpcserverinit "ecs_govel/internal/grpcservice/app/server"
	"ecs_govel/internal/restroute"

	"ecs_govel/pkg/pkggin"
	"ecs_govel/pkg/pkglog"
	"fmt"
	"log"

	_ "ecs_govel/pkg/pkgmetrics/sdkopentelemetry"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func run() {
	log.Printf("%v", pkglog.Log)
	go command.RunJobs()

	// Init GrpcService
	if strings.ToLower(config.TYPE_SERVICES) == "grpc" {
		restroute.LoadInfoHandlers()

		go pkggin.HttpRun()

		fmt.Println("InitGrpcService")
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
