package config

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var Database = new(DbInstance)

type DbInstance struct {
	*gorm.DB
	err   error
	AppID string
}

func init() {
	godotenv.Load()
	driver := strings.ToLower(os.Getenv("DB_DRIVER"))
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	db := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASSWD")
	ssl := os.Getenv("DB_SSL_MODE")
	ConnStr := ""
	switch driver {
	case "postgres":
		ConnStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			host, port, user, db, pass, ssl)
		break
	case "mysql":
		ConnStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			host, port, user, db, pass, ssl)
		break
	}

	log.Println(ConnStr)

	var err error
	Database.DB, err = gorm.Open(driver, ConnStr)
	Database.err = err
	if err != nil {
		log.Println("Ocurrio um error", err)
	}
	log.Printf("New %s conennection opened\n", driver)
	Database.Init()
}

func (db *DbInstance) GetDbInstance() *gorm.DB {
	return Database.DB
}

func (db *DbInstance) Init() {
	//init1()
	//db.DB = db.GetDbInstance()
	if db.err != nil {
		log.Println(db.err)
	}

	//Configiuracion de conecciones idle
	db.DB.DB().SetConnMaxLifetime(30 * time.Minute)
	db.DB.DB().SetMaxIdleConns(100)
	db.DB.DB().SetMaxOpenConns(100)

	rand.Seed(time.Now().UnixNano())

	// Create tables if they dont exist or migrate schema if exist
	// if !db.DB.HasTable(&migrations.User{}) {
	// 	db.DB.CreateTable(&migrations.User{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.User{})
	// }

	// if !db.DB.HasTable(&migrations.Application{}) {
	// 	db.DB.CreateTable(&migrations.Application{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.Application{})
	// }

	// if !db.DB.HasTable(&migrations.WhatchDog{}) {
	// 	db.DB.CreateTable(&migrations.WhatchDog{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.WhatchDog{})
	// }

	// if !db.DB.HasTable(&migrations.WPAccounts{}) {
	// 	db.DB.CreateTable(&migrations.WPAccounts{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.WPAccounts{})
	// }

	// if !db.DB.HasTable(&migrations.History{}) {
	// 	db.DB.CreateTable(&migrations.History{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.History{})
	// }

	// if !db.DB.HasTable(&migrations.Sessions{}) {
	// 	db.DB.CreateTable(&migrations.Sessions{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.Sessions{})
	// }

	// if !db.DB.HasTable(&migrations.CompanyWhatsapps{}) {
	// 	db.DB.CreateTable(&migrations.CompanyWhatsapps{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.CompanyWhatsapps{})
	// }

	// if !db.DB.HasTable(&migrations.Chats{}) {
	// 	db.DB.CreateTable(&migrations.Chats{})
	// } else {
	// 	db.DB.AutoMigrate(&migrations.Chats{})
	// }
}
