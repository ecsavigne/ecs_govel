package config

import (
	"database/sql"
	"ecs_govel/app/helpers/logg"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	//"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type migration map[string]string

type Migrate struct {
	isAuto bool
	// mapa de nombre de migraciones que almacena las path /database/migrations/*.sql
	// Asociada a la migrationName Key del mapa
	migrations migration
	// maneja el tiempo en que son cargadas las migrations
	timeFirstLoad time.Time
}
type DbInstance struct {
	*gorm.DB
	//Obj que tiene informacion realacionada con las migration
	Migrates Migrate
	AppID    string
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
	fmt.Printf("Cargando info Base Datos de: %s\n", pathEnv)
	err_ := godotenv.Load(pathEnv)
	if err_ != nil {
		logg.ErrorLogger.Printf("Error: \033[31m%v\033[0m\n", err_)
		fmt.Println("Error:Error: \033[31m cargando Var ambiente: ", err_.Error(), "\033[0m")
		return
	}
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
		fmt.Println("driver: POSTGRESS")
		ConnStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=Asia/Shanghais",
			host, port, user, db, pass, ssl)
		Database.DB, err = gorm.Open(postgres.Open(ConnStr), &gorm.Config{})
		break
	case "mysql":
		fmt.Println("driver: MYSQL")
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

	Database.Migrates.isAuto = configsIni.autoMigration

	logg.GeneralLogger.Printf("New (%s) conennection opened\n", driver)
	fmt.Printf("New (%s) conennection opened\n", driver)

	logg.GeneralLogger.Printf("\033[36mConfigurando coneccion y cargando Migration %s: \033[0m\n", driver)
	fmt.Printf("\033[36mConfigurando coneccion y cargando Migration de:  %s \033[0m\n", configsIni.folderMigrations)
	Database.loadMigrationNameFromMigrationsFolder()
	Database.autoMigrate()
}

func (db *DbInstance) db() *sql.DB {
	DB, _ := db.DB.DB()
	return DB
}

// Esta funcion permite ejecutar las migraciones cuando se carga la instancia
// de DbInstance, si DbInstance.Migrate.isAuto == true sino no carga
func (db *DbInstance) autoMigrate() {
	if db.Migrates.isAuto {
		db.execAllMigration()
		db.Migrates.timeFirstLoad = time.Now()
		fmt.Println("Activado carga migration automaticas")
	} else {
		fmt.Println("Desactivado carga migration automaticas")
	}
	//db.execOneMigration("Otra")
}

// Esta function carga todos los nombres de las migration segun el nombre del
// archivo .go situado en la carpeta que se almacenaran las migration y la asocia a su
// respectivo archivo .sql
func (db *DbInstance) loadMigrationNameFromMigrationsFolder() {
	db.Migrates.migrations = make(migrations)
	files, err := os.ReadDir(configsIni.folderMigrations)
	if err != nil {
		logg.ErrorLogger.Printf("Error \033[31mal leer la carpeta Error: %+v\033[0m\n", err)
		fmt.Println("Error \033[31mal leer la carpeta:", err, "\033[0m")
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == ".sql" {
			strTemp := deletePatronOffString(strings.TrimSuffix(file.Name(), ".sql"))
			if strTemp != "" {
				db.Migrates.migrations[strTemp] = configsIni.folderMigrations + file.Name()
			}
		}
	}

	fmt.Println("Folder Migrate:", configsIni.folderMigrations)
	fmt.Printf("Migrates:\n\t%+v\n", db.Migrates.migrations)
}

func (db *DbInstance) initDB() {
	// Configiuracion de conecciones idle
	db.db().SetConnMaxLifetime(time.Duration(configsIni.setConnMaxLifetime) * time.Minute)
	db.db().SetMaxIdleConns(configsIni.setMaxIdleConns)
	db.db().SetMaxOpenConns(configsIni.setMaxOpenConns)

	// Cargar de Migration
	//rand.Seed(time.Now().UnixNano())

}

// Ejecutar migration dado un nombre
func (db *DbInstance) execOneMigration(nameSql string) {
	file, err := os.Open(db.Migrates.migrations[nameSql])
	if err != nil {
		logg.GeneralLogger.Printf("fallo al abrir file:%s,  Error: %v \n", nameSql, err)
		fmt.Printf("fallo al abrir file:%s,  Error: %v \n", nameSql, err)
		return
	}
	defer file.Close()

	queries, err := io.ReadAll(file)
	if err != nil {
		logg.GeneralLogger.Printf("fallo al leer file:%s,  Error: %v \n", nameSql, err)
		fmt.Printf("fallo al leer file:%s,  Error: %v \n", nameSql, err)
		return
	}

	// Ejecuta las consultas SQL
	db.Exec(string(queries))
	if db.Error != nil {
		logg.GeneralLogger.Printf("Ocurrio un fallo ejecutando Query Error:%v \n", err)
		fmt.Printf("Ocurrio un fallo ejecutando Query Error:%v \n", err)
		return
	}
}

// Carga todas las migration situada en el directorio de migration
func (db *DbInstance) execAllMigration() {
	for _, v := range db.Migrates.migrations {
		db.execOneMigration(db.Migrates.migrations[v])
	}
}

// Cargara grupo de migration especificadas que deben estar en el dir de migrations
func (db *DbInstance) ExecSetMigration(nombreMigration []string) {
	for i := range nombreMigration {
		db.execOneMigration(nombreMigration[i])
	}
}
