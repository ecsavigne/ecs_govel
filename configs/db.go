package configs

import (
	model "ecs_govel/app/model"
	migrations "ecs_govel/database/migration"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	// model "ecs_govel/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"oficial_gin/app/model"
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
			ErrorLogger.Println("[database.database.go - init()]. ", logMessage)
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

	model.SetGlobalDB(Database)
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
			ErrorLogger.Println("[database.database.go - Init()]. ", logMessage)
			fmt.Println(logMessage)
		}
	}()

	debugMessage = "2"
	//db.DB = db.GetDbInstance()
	// if db.err != nil {
	// 	panic(db.err)
	// }

	// Create tables if they dont exist and migrate schema if exist
	// if !db.DB.Migrator().HasTable(&model.User{}) {
	// 	db.DB.Migrator().CreateTable(&model.User{})
	// } else {
	// 	db.DB.AutoMigrate(&model.User{})
	// }

	if !db.DB.Migrator().HasTable(&migrations.Application{}) {
		db.DB.Migrator().CreateTable(&migrations.Application{})
	} else {
		db.DB.AutoMigrate(&migrations.Application{})
		var v = migrations.Application{}
		db.DB.Model(&v).
			Preload("WPAccounts").
			Order("id ASC").
			Find(&v)
		if c := len(v.WPAccounts); c != 0 {
			fmt.Printf("Existen \033[34m%d\033[0m activas WPAccounts asociados a la aplicacion\n", c)
			// fmt.Println("Cantidad de cuentas asosciadas a la aplicacion: ", c)
			// jsoN, _ := json.MarshalIndent(v, "", " ")
			// fmt.Printf("AplicationData:\n", string(jsoN), "\n")
		} else {
			fmt.Println("No existen WPAccounts asociados a la aplicacion")
		}
	}

	// if !db.DB.Migrator().HasTable(&migrations.WhatchDog{}) {
	// 	db.DB.Migrator().CreateTable(&migrations.WhatchDog{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.WhatchDog{})
	// }

	// if !db.DB.Migrator().HasTable(&migrations.WPAccounts{}) {
	// 	db.DB.Migrator().CreateTable(&migrations.WPAccounts{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.WPAccounts{})
	// }

	// if !db.DB.Migrator().HasTable(&migrations.History{}) {
	// 	db.DB.Migrator().CreateTable(&migrations.History{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.History{})
	// }

	// if !db.DB.Migrator().HasTable(&migrations.Sessions{}) {
	// 	db.DB.Migrator().CreateTable(&migrations.Sessions{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.Sessions{})
	// }

	// if !db.DB.Migrator().HasTable(&migrations.CompanyWhatsapps{}) {
	// 	db.DB.Migrator().CreateTable(&migrations.CompanyWhatsapps{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.CompanyWhatsapps{})
	// }

	// Nos se puede correr porque tiene mucha informacion y estresa el servidor
	// if !db.DB.Migrator().HasTable(&migrations.Chats{}) {
	// 	db.DB.Migrator().CreateTable(&migrations.Chats{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.Chats{})
	// }

	// if !db.DB.Migrator().HasTable(&migrations.ShippingOpenRate{}) {
	// 	db.DB.Migrator().CreateTable(&migrations.ShippingOpenRate{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.ShippingOpenRate{})
	// }

	// if !db.DB.Migrator().HasTable(&migrations.Color{}) {
	// 	db.DB.Migrator().CreateTable(&migrations.Color{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.Color{})
	// }

	// if !db.DB.Migrator().HasTable(&migrations.Tag{}) {
	// 	db.DB.Migrator().CreateTable(&migrations.Tag{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.Tag{})
	// }

	// if !db.DB.Migrator().HasTable(&migrations.Contact{}) {
	// 	fmt.Println("bd .5")
	// 	db.DB.Migrator().CreateTable(&migrations.Contact{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.Contact{})
	// }

	// if !db.DB.Migrator().HasTable(&migrations.ChatTag{}) {

	// 	db.DB.Migrator().CreateTable(&migrations.ChatTag{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.ChatTag{})
	// }

	// if !db.DB.Migrator().HasTable(&migrations.CompanyWhatsappTag{}) {
	// 	db.DB.Migrator().CreateTable(&migrations.CompanyWhatsappTag{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.CompanyWhatsappTag{})
	// }

	// if !db.DB.Migrator().HasTable(&migrations.ContactTag{}) {
	// 	db.DB.Migrator().CreateTable(&migrations.ContactTag{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.ContactTag{})
	// }

}

func (db *DbInstance) CreateApp() {
	application := model.Application{
		AppName:     "socialhub_chat",
		APIHash:     "S4h_EPRZm-b46kyoUbUJ",
		WebhookText: "http://principal.socialhub.local/reciveTextMessage",
		WebhookFile: "http://principal.socialhub.local/reciveFileMessage",
	}

	//TODO db.DB.Create(&application) intentaba criar y si ya existia el registro devolvia error de primary key  ERROR: duplicate key value violates unique constraint "applications_api_hash_key"
	if res := db.DB.FirstOrCreate(&application); res.Error != nil {
		fmt.Println("Error creating application ", application, ". Error is: ", res.Error.Error())
		return
	}
}
