package config

import (
	"database/sql"
	"ecs_govel/app/helpers/logg"
	"fmt"
	"os"
	"regexp"
	"strings"

	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	*gorm.DB
	//err   error
	AppID string
}

var Database = new(DbInstance)

/*
Inicializa configuracion de la base de datos
Se le pasa '@pathEnv'=> Dir del fichero .env de las configuraciones de BD
'@driverP' => Tipo de driver de base de datos
*/
func configDB(pathEnv, driverP string) {
	if driverP == "" {
		logg.ErrorLogger.Printf("Error: \033[31m%v\033[0m\n", "Driver BD no presente")
		fmt.Println("Driver BD no presente")
		return
	}
	logg.GeneralLogger.Printf("Cargando info Base Datos\n")
	fmt.Printf("Cargando info Base Datos de %s\n", pathEnv)
	err_ := godotenv.Load(pathEnv)
	if err_ != nil {
		logg.ErrorLogger.Printf("Error: \033[31m%v\033[0m\n", err_)
		fmt.Println("Error:Error: \033[31m cargando Var ambiente: ", err_.Error(), "\033[0m")
		return
	}
	fmt.Println("111111111", driverP)
	driver := strings.ToLower(driverP)
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	db := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASSWD")
	ssl := os.Getenv("DB_SSL_MODE")
	ConnStr := ""
	var err error
	switch driver {
	case "postgres":
		fmt.Println("POSTGRESS")
		ConnStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=Asia/Shanghais",
			host, port, user, db, pass, ssl)
		Database.DB, err = gorm.Open(postgres.Open(ConnStr), &gorm.Config{})
		break
	case "mysql":
		fmt.Println("MYSQL")
		ConnStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, pass, host, port, db)
		Database.DB, err = gorm.Open(mysql.Open(ConnStr), &gorm.Config{})
		break
	default:
		logg.ErrorLogger.Printf("Error: \033[31m%v\033[0m\n", "Driver de BD no identificado")
		fmt.Println("Driver de BD no identificado")
		return
		break
	}
	if err != nil {
		expresionRegular := regexp.MustCompile("")
		if expresionRegular.MatchString(err.Error()) {
			fmt.Printf("Error: \033[31mVerificar que la configuracion sea la adecuada para base datos tipo: (%s)\033[0m .\n", driverP)
			logg.ErrorLogger.Printf("Error: \033[31mVerificar que la configuracion se la adecuada para base datos tipo: \033[31m(%s)\033[0m .\n", driverP)
		} else {
			fmt.Println("Error", err)
			logg.ErrorLogger.Println("Ocurrio um error: \033[31m %s\033[0m", err)
		}
		return
	}

	logg.GeneralLogger.Printf("New %s conennection opened\n", driver)
	// Database.initDB()
}

func (db *DbInstance) db() *sql.DB {
	DB, _ := db.DB.DB()
	return DB
}

func (db *DbInstance) GetDbInstance() *gorm.DB {
	return Database.DB
}

func (db *DbInstance) initDB() {
	// //init1()
	// //db.DB = db.GetDbInstance()
	// if db.err != nil {
	// 	log.Println(db.err)
	// }

	// //Configiuracion de conecciones idle
	// db.DB.DB().SetConnMaxLifetime(30 * time.Minute)
	//db.DB.DB().SetMaxIdleConns(100)
	//db.DB.DB().SetMaxOpenConns(100)

	//rand.Seed(time.Now().UnixNano())

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
