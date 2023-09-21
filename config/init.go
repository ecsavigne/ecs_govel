package config

import (
	"ecs_govel/app/helpers/logg"
	"fmt"
	"log"
	"strings"

	"github.com/go-ini/ini"
)

type configEcsGovelIni struct {
	//Files .ini
	cfgAppIni *ini.File
	cfgBDIni  *ini.File
	//Seccion de variables de App
	sessionApp       *ini.Section
	appID            string
	folderMigrations string
	autoMigration    bool
	//Seccion de variables de Servidor Web
	sessionWeb   *ini.Section
	host         string
	port         int
	writeTimeout int
	readTimeout  int
	idleTimeout  int
	//Seccion de variables de Base de datos
	sessionDB          *ini.Section
	urlEnv             string
	driver             string
	setConnMaxLifetime int
	setMaxIdleConns    int
	setMaxOpenConns    int
}

var configsIni = configEcsGovelIni{}

func init() {
	var errConfig error
	configsIni.cfgAppIni, errConfig = ini.Load("./config/app.ini")
	if errConfig != nil {
		log.Fatalf("Error al cargar el archivo .ini: %v", errConfig)
	}
	configsIni.cfgBDIni, errConfig = ini.Load("./config/db.ini")
	if errConfig != nil {
		log.Fatalf("Error al cargar el archivo .ini: %v", errConfig)
	}

	configsIni.sessionApp = configsIni.cfgAppIni.Section("App")
	configsIni.appID, Database.AppID = configsIni.sessionApp.Key("AppID").String(), configsIni.sessionApp.Key("AppID").String()
	configsIni.folderMigrations = configsIni.sessionApp.Key("FolderMigrations").String()
	configsIni.autoMigration, _ = configsIni.sessionApp.Key("AutoMigration").Bool()
	configsIni.sessionWeb = configsIni.cfgAppIni.Section("Servidor Web")
	configsIni.host = configsIni.sessionWeb.Key("Host").String()
	configsIni.port, _ = configsIni.sessionWeb.Key("Port").Int()
	configsIni.writeTimeout, _ = configsIni.sessionWeb.Key("WriteTimeout").Int()
	configsIni.writeTimeout, _ = configsIni.sessionWeb.Key("ReadTimeout").Int()
	configsIni.idleTimeout, _ = configsIni.sessionWeb.Key("IdleTimeout").Int()
	configsIni.sessionDB = configsIni.cfgBDIni.Section("Env DataBase")
	configsIni.driver = configsIni.sessionDB.Key("Driver").String()
	configsIni.urlEnv = configsIni.sessionDB.Key("UrlEnv").String()
	configsIni.setConnMaxLifetime, _ = configsIni.sessionDB.Key("SetConnMaxLifetime").Int()
	configsIni.setMaxIdleConns, _ = configsIni.sessionDB.Key("SetMaxIdleConns").Int()
	configsIni.setMaxOpenConns, _ = configsIni.sessionDB.Key("SetMaxOpenConns").Int()
}

// Elimina un patron de una cadena y retorna desde el inicio hasta el patron
// Se usa para eliminar Migration.sql de los archivos  que estan en la ruta de las
// Migration seguen el criterio aplicado
func deletePatronOffString(cad string) string {
	patron := "Migration"
	indice := strings.Index(cad, patron)
	if indice == -1 {
		logg.ErrorLogger.Println("Error: \033[31Patron no existente: %\033[0m\n")
		fmt.Println("Error: \033[31Patron no existente: %\033[0m\n")
		return ""
	}
	strTemp := cad[:indice]
	return strings.ToUpper(string(strTemp[0])) + strTemp[1:]
}
