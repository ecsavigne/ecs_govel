package config

import (
	"ecs_govel/app/helpers/logg"
	//"ecs_govel/config"
	"ecs_govel/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-ini/ini"
)

var cfgAppIni *ini.File
var cfgBDIni *ini.File
var errConfig error
var sessionApp *ini.Section
var sessionServerWeb *ini.Section
var sessionDataBase *ini.Section

func init() {
	cfgAppIni, errConfig = ini.Load("./config/app.ini")
	if errConfig != nil {
		log.Fatalf("Error al cargar el archivo .ini: %v", errConfig)
	}
	cfgBDIni, errConfig = ini.Load("./config/app.ini")
	if errConfig != nil {
		log.Fatalf("Error al cargar el archivo .ini: %v", errConfig)
	}

	sessionDataBase = cfgBDIni.Section("Env DataBase")
	sessionApp = cfgAppIni.Section("App")
	sessionServerWeb = cfgAppIni.Section("Servidor Web")
	fmt.Println("Cargo VAr .ini")
	configDB(sessionDataBase.Key("UrlEnv").String())
}

func InitApp() {
	fmt.Println(Database.AppID)
	// MAnejo de errores
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ocurrio algun error")
		}
	}()

	// http server configuration for this main process
	writeTime, _ := sessionServerWeb.Key("WriteTimeout").Int()
	readTime, _ := sessionServerWeb.Key("ReadTimeout").Int()
	idleTime, _ := sessionServerWeb.Key("IdleTimeout").Int()
	server := &http.Server{
		Handler:      routes.Router,
		Addr:         sessionServerWeb.Key("Host").String() + ":" + sessionServerWeb.Key("Port").String(),
		WriteTimeout: time.Second * time.Duration(writeTime),
		ReadTimeout:  time.Second * time.Duration(readTime),
		IdleTimeout:  time.Second * time.Duration(idleTime),
	}

	// 10. run the http server paralelly in a goroutine to receive request
	go func() {
		logg.GeneralLogger.Printf("Iniciando server --- %s:%s\n", sessionServerWeb.Key("Host").String(), sessionServerWeb.Key("Port").String())
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
