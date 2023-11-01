package config

import (
	"ecs_govel/app/helpers/logg"
	"ecs_govel/config/db"
	"ecs_govel/routes"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitApp() {
	fmt.Println("Id App:", db.Orm.AppID)
	//db.ConfigDB(configsIni.urlEnv, configsIni.driver)
	db.ConfigDB(configsIni)
	// MAnejo de errores
	defer func() {
		if err := recover(); err != nil {
			logg.GeneralLogger.Println("Ocurrio algun error:", err)
			fmt.Println("Ocurrio algun error:", err)
		}
	}()

	server := &http.Server{
		Handler:      routes.Router,
		Addr:         fmt.Sprintf("%s:%d", configsIni.Host, configsIni.Port),
		WriteTimeout: time.Second * time.Duration(configsIni.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(configsIni.ReadTimeout),
		IdleTimeout:  time.Second * time.Duration(configsIni.IdleTimeout),
	}
	fmt.Printf("Sever Web en: %s:%d\n", configsIni.Host, configsIni.Port)
	// Run the http server paralelly in a goroutine to receive request
	go func() {
		logg.GeneralLogger.Printf("Iniciando server --- %s:%d\n", configsIni.Host, configsIni.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logMessage := "Impossible initialice server: " + err.Error()
			logg.GeneralLogger.Println(logMessage)
			return
		}
	}()
	// Espera por Ctrl+C para finalizar la aplicacion
	c := make(chan os.Signal, 1)                    //registra el canal c para recibir la señal SIGINT (representada por os.Interrupt).
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // La aplicación espere hasta que reciba la señal SIGINT.
	<-c
	server.Close()
	logg.GeneralLogger.Println("Se preciono Ctrl + C, para finalizar la App")
}
