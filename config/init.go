package config

import (
	"log"

	"ecs_govel/config/db"

	"github.com/go-ini/ini"
)

type Config interface {
	GetObj() interface{}
}

// type ConfigEcsGovelIni struct {
// 	//Files .ini
// 	cfgAppIni *ini.File
// 	cfgBDIni  *ini.File
// 	//Seccion de variables de App
// 	sessionApp         *ini.Section
// 	appID              string
// 	folderMigrations   string
// 	autoMigration      bool
// 	migrationFromModel bool
// 	//Seccion de variables de Servidor Web
// 	sessionWeb   *ini.Section
// 	host         string
// 	port         int
// 	writeTimeout int
// 	readTimeout  int
// 	idleTimeout  int
// 	//Seccion de variables de Base de datos
// 	sessionDB          *ini.Section
// 	UrlEnv             string
// 	Driver             string
// 	setConnMaxLifetime int
// 	setMaxIdleConns    int
// 	setMaxOpenConns    int
// 	activeOnCascade    bool
// }

// func (conf ConfigEcsGovelIni) GetObj() interface{} {
// 	return conf
// }

var configsIni = db.ConfigEcsGovelIni{}

func init() {
	var errConfig error
	configsIni.CfgAppIni, errConfig = ini.Load("./config/app.ini")
	if errConfig != nil {
		log.Fatalf("Error al cargar el archivo .ini: %v", errConfig)
	}
	configsIni.CfgBDIni, errConfig = ini.Load("./config/db/db.ini")
	if errConfig != nil {
		log.Fatalf("Error al cargar el archivo .ini: %v", errConfig)
	}

	configsIni.SessionApp = configsIni.CfgAppIni.Section("App")
	configsIni.AppID, db.Orm.AppID = configsIni.SessionApp.Key("AppID").String(), configsIni.SessionApp.Key("AppID").String()
	configsIni.FolderMigrations = configsIni.SessionApp.Key("FolderMigrations").String()
	configsIni.AutoMigration, _ = configsIni.SessionApp.Key("AutoMigration").Bool()
	configsIni.MigrationFromModel, _ = configsIni.SessionApp.Key("MigrationFromModel").Bool()
	configsIni.SessionWeb = configsIni.CfgAppIni.Section("Servidor Web")
	configsIni.Host = configsIni.SessionWeb.Key("Host").String()
	configsIni.Port, _ = configsIni.SessionWeb.Key("Port").Int()
	configsIni.WriteTimeout, _ = configsIni.SessionWeb.Key("WriteTimeout").Int()
	configsIni.WriteTimeout, _ = configsIni.SessionWeb.Key("ReadTimeout").Int()
	configsIni.IdleTimeout, _ = configsIni.SessionWeb.Key("IdleTimeout").Int()
	configsIni.SessionDB = configsIni.CfgBDIni.Section("Env DataBase")
	configsIni.Driver = configsIni.SessionDB.Key("Driver").String()
	configsIni.UrlEnv = configsIni.SessionDB.Key("UrlEnv").String()
	configsIni.ActiveOnCascade, _ = configsIni.SessionDB.Key("ActiveOnCascade").Bool()
	configsIni.SetConnMaxLifetime, _ = configsIni.SessionDB.Key("SetConnMaxLifetime").Int()
	configsIni.SetMaxIdleConns, _ = configsIni.SessionDB.Key("SetMaxIdleConns").Int()
	configsIni.SetMaxOpenConns, _ = configsIni.SessionDB.Key("SetMaxOpenConns").Int()

}
