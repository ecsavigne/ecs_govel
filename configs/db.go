package configs

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"ecs_govel/app/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	// DB    *gorm.DB
	*gorm.DB
	AppID string
}

var (
	DB_HOST         string
	DB_USER         string
	DB_NAME         string
	DB_PASSWORD     string
	DB_PORT         string
	FORWARD_DB_PORT string
	DB_CONNSTR      string
	DNS_DB          string
	Database        *DbInstance = new(DbInstance)
)

func create_database() {}

func prepare_db() {
	var (
		logMessage, debugMessage string
		err                      error
	)

	defer func() {
		if r := recover(); r != nil {
			logMessage = filepath.Base(os.Args[0]) + " :  Error initializing Database. " + ": Recovered from exception " + ". Interface in defer is: " + fmt.Sprintf("%+v", r) + ". DebugMessage is: " + debugMessage
			Log.Infof("[database.database.go - init()]. ", logMessage)
			os.Exit(2)
		}
	}()

	// Create database si no existe
	create_database()

	DB_CONNSTR = fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", DB_HOST, DB_USER, DB_NAME, FORWARD_DB_PORT, DB_PASSWORD)

	Database.DB, err = gorm.Open(postgres.Open(DB_CONNSTR), &gorm.Config{})
	dateFormat := time.Now()
	now := dateFormat.Format("2006-01-02 15:04:05")
	fmt.Println("New postgres conennection opened at ", now)
	if err != nil {
		panic(err)
	}

	model.SetGlobalDB(Database.DB)
	fmt.Println("Max Connections: ", APP_MAX_CONNECTIONS, " CantX: ", APP_CANT_X)
	sqlDB, _ := Database.DB.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 2) // Make than last forever
	sqlDB.SetMaxIdleConns((APP_MAX_CONNECTIONS / APP_CANT_X) - 7)
	sqlDB.SetMaxOpenConns((APP_MAX_CONNECTIONS / APP_CANT_X) - 5)

	rand.New(rand.NewSource(time.Now().UnixNano()))
	// Agregar plugin para prometeus
	//fmt.Println("aAse datos mericas:", fmt.Sprintf("%s:%s", HTTP_SERVER_HOST_METRICS, HTTP_SERVER_PORT_METRICS))
	// Database.DB.Use(prometheus.New(prometheus.Config{
	// 	DBName:          DB_NAME,
	// 	RefreshInterval: 15,
	// 	PushAddr:        fmt.Sprintf("%s:%s", HTTP_SERVER_HOST_METRICS, HTTP_SERVER_PORT_METRICS),
	// 	StartServer:     false,
	// 	//MetricsCollector: []prometheus.MetricsCollector{},
	// }))

	// Load Migration
	Database.Migrate()
}

// Init : First Setup
func (db *DbInstance) Migrate() {
	logMessage := ""
	debugMessage := "1"
	defer func() {
		if r := recover(); r != nil {
			logMessage = filepath.Base(os.Args[0]) + " :  Error initializing Migration in Database. " + ": Recovered from exception " + ". Interface in defer is: " + fmt.Sprintf("%+v", r) + ". DebugMessage is: " + debugMessage
			Log.Infof("[database.database.go - Init()]. ", logMessage)
			fmt.Println(logMessage)
		}
	}()

	// if !db.DB.Migrator().HasTable(&migrations.WhatchDog{}) {
	// 	db.DB.Migrator().CreateTable(&migrations.WhatchDog{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.WhatchDog{})
	// }
}
