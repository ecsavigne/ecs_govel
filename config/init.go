package config

import (
	"log"

	"github.com/go-ini/ini"
)

type configEcsGovelIni struct {
	//Files .ini
	cfgAppIni *ini.File
	cfgBDIni  *ini.File
	//Seccion de variables de App
	sessionApp *ini.Section
	appID      string
	//Seccion de variables de Servidor Web
	sessionWeb   *ini.Section
	host         string
	port         int
	writeTimeout int
	readTimeout  int
	idleTimeout  int
	//Seccion de variables de Base de datos
	sessionDB *ini.Section
	urlEnv    string
	driver    string
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
	Database.AppID = configsIni.sessionApp.Key("AppID").String()
	configsIni.sessionWeb = configsIni.cfgAppIni.Section("Servidor Web")
	configsIni.writeTimeout, _ = configsIni.sessionWeb.Key("WriteTimeout").Int()
	configsIni.writeTimeout, _ = configsIni.sessionWeb.Key("ReadTimeout").Int()
	configsIni.idleTimeout, _ = configsIni.sessionWeb.Key("IdleTimeout").Int()
	configsIni.sessionDB = configsIni.cfgBDIni.Section("Env DataBase")
	configsIni.driver = configsIni.sessionDB.Key("Driver").String()
	configsIni.urlEnv = configsIni.sessionDB.Key("UrlEnv").String()
	//configDB(sessionDataBase.Key("UrlEnv").String(), sessionDataBase.Key("Driver").String())
}
