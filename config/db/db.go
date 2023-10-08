package db

import (
	"database/sql"
	"ecs_govel/app/helpers/logg"
	"ecs_govel/app/models"

	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/go-ini/ini"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigEcsGovelIni struct {
	//Files .ini
	CfgAppIni *ini.File
	CfgBDIni  *ini.File
	//Seccion de variables de App
	SessionApp         *ini.Section
	AppID              string
	FolderMigrations   string
	AutoMigration      bool
	MigrationFromModel bool
	//Seccion de variables de Servidor Web
	SessionWeb   *ini.Section
	Host         string
	Port         int
	WriteTimeout int
	ReadTimeout  int
	IdleTimeout  int
	//Seccion de variables de Base de datos
	SessionDB          *ini.Section
	UrlEnv             string
	Driver             string
	SetConnMaxLifetime int
	SetMaxIdleConns    int
	SetMaxOpenConns    int
	ActiveOnCascade    bool
}

var configsIni ConfigEcsGovelIni

type migration map[string]string
type migrationPath struct {
	key  string
	path string
}
type migrationOrder []migrationPath

type Migrate struct {
	isAuto               bool
	isMigrationFromModel bool
	// mapa de nombre de migraciones que almacena las path /database/migrations/*.sql
	// Asociada a la migrationName Key del mapa
	migrations migration
	// Migraciones ordenadas por Key
	migrationsOrders migrationOrder
	// maneja el tiempo en que son cargadas las migrations
	timeFirstLoad time.Time
}
type DbInstance struct {
	*gorm.DB
	//Obj que tiene informacion realacionada con las migration
	Migrates Migrate
	AppID    string
}

var Orm = new(DbInstance)

/*
Inicializa configuracion de la base de datos
Se le pasa '@pathEnv'=> Dir del fichero .env de las configuraciones de BD
'@driverP' => Tipo de driver de base de datos
*/
//func ConfigDB(pathEnv, driverP string) {
func ConfigDB(obj ConfigEcsGovelIni) {
	configsIni = obj
	driverP := obj.Driver
	pathEnv := obj.UrlEnv
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
		fmt.Println("Error: \033[31m cargando Var ambiente: ", err_.Error(), "\033[0m")
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
		// DisableForeignKeyConstraintWhenMigrating: false activa la actualizacion en cascada cuando migra
		// hay que correr las migration con true y luego con false para agregar actualizacion
		// en cascada o ordenar las tablas segun el orden que acomoda o sino la otra opcion
		// de cargar migration es decir que no sea via gorm sino .sql
		Orm.DB, err = gorm.Open(postgres.Open(ConnStr), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: !configsIni.ActiveOnCascade})
		break
	case "mysql":
		fmt.Println("driver: MYSQL")
		ConnStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, pass, host, port, db)
		// DisableForeignKeyConstraintWhenMigrating: false activa la actualizacion en cascada cuando migra
		Orm.DB, err = gorm.Open(mysql.Open(ConnStr), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: !configsIni.ActiveOnCascade})
		break
	default:
		logg.ErrorLogger.Printf("Error: \033[31m%v\033[0m\n", "Driver de BD no identificado")
		fmt.Println("Driver de BD no identificado")
		break
	}
	if err != nil {
		expresionRegularError := regexp.MustCompile("")
		expRegServerDBOff := regexp.MustCompile(`dial tcp \[::`)
		expRegErrorUrlDB := regexp.MustCompile(`invalid URL escape`)
		if expresionRegularError.MatchString(err.Error()) {
			if expRegServerDBOff.MatchString(err.Error()) {
				fmt.Printf("Error: \033[31mServidor de Base de datos tipo: (%s) esta off. Inicielo!!!!!\033[0m .\n", driverP)
				logg.ErrorLogger.Printf("Error: \033[31mServidor de Base de datos tipo: (%s) esta off. Inicielo!!!!!\033[0m .\n", driverP)
			}
			if expRegErrorUrlDB.MatchString(err.Error()) {
				fmt.Printf("Error: \033[31mConfiguracion de Base de datos: (%s) incompleta. Verifiquela!!!!!\033[0m .\n config URL: (%s)", driverP, ConnStr)
				logg.ErrorLogger.Printf("Error: \033[31mConfiguracion de Base de datos: (%s) incompleta. Verifiquela!!!!!\033[0m .\n config URL: (%s)", driverP, ConnStr)
			} else {
				fmt.Printf("Error: \033[31mVerificar que la configuracion sea la adecuada para base datos tipo: (%s)\033[0m .\n", driverP)
				logg.ErrorLogger.Printf("Error: \033[31mVerificar que la configuracion se la adecuada para base datos tipo: \033[31m(%s)\033[0m .\n", driverP)
			}
		} else {
			fmt.Println("Ocurrio um error: \033[31m %s\033[0m", err)
			logg.ErrorLogger.Println("Ocurrio um error: \033[31m %s\033[0m", err)
		}
		return
	}

	Orm.Migrates.isAuto = configsIni.AutoMigration
	Orm.Migrates.isMigrationFromModel = configsIni.MigrationFromModel

	logg.GeneralLogger.Printf("New (%s) conennection opened\n", driver)
	fmt.Printf("New (%s) conennection opened\n", driver)

	logg.GeneralLogger.Printf("\033[36mConfigurando coneccion y cargando Migration %s: \033[0m\n", driver)
	fmt.Printf("\033[36mConfigurando coneccion y cargando Migration de:  %s \033[0m\n", configsIni.FolderMigrations)
	Orm.autoMigrate()
}

func (db *DbInstance) db() *sql.DB {
	DB, _ := db.DB.DB()
	return DB
}

// Copia en db.Migrates.migrationsOrder db.Migrates.migrations
func (db *DbInstance) orderMigrationPath() {
	sort.Slice(db.Migrates.migrationsOrders, func(i, j int) bool {
		return db.Migrates.migrationsOrders[i].key < db.Migrates.migrationsOrders[j].key
	})
	//fmt.Printf("MIGRA ORDER:%v\n", db.Migrates.migrationsOrders)
}

// Esta funcion permite ejecutar las migraciones cuando se carga la instancia
// de DbInstance, si DbInstance.Migrate.isAuto == true sino no carga
func (db *DbInstance) autoMigrate() {
	Orm.loadMigrationNameFromMigrationsFolder()
	if db.Migrates.isAuto == true && db.Migrates.isMigrationFromModel == true {
		fmt.Println("\033[31mNo puede estar los 2 modos de carga de migration automatica activa, revise app.ini en config/app.ini\033[0m")
		logg.GeneralLogger.Println("\033[31mNo puede estar los 2 modos de carga de migration automatica activa, revise app.ini en config/app.ini\033[0m")
		return
	}
	if db.Migrates.isAuto == true {
		db.execAllMigration()
		db.Migrates.timeFirstLoad = time.Now()
		fmt.Println("\033[31mActivado carga migration automaticas\033[0m")
		logg.GeneralLogger.Println("\033[31mActivado carga migration automaticas\033[0m")
	} else if db.Migrates.isMigrationFromModel == true {
		db.execAllMigrationFromModels()
	} else {
		fmt.Println("\033[36mDesactivado carga migration automaticas\033[0m")
		logg.GeneralLogger.Println("\033[36mDesactivado carga migration automaticas\033[0m")
	}
}

// Esta function carga todos los nombres de las migration segun el nombre del
// archivo .go situado en la carpeta que se almacenaran las migration y la asocia a su
// respectivo archivo .sql
func (db *DbInstance) loadMigrationNameFromMigrationsFolder() {
	db.Migrates.migrations = make(migration)
	db.Migrates.migrationsOrders = make(migrationOrder, 0)
	files, err := os.ReadDir(configsIni.FolderMigrations)
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
				db.Migrates.migrations[strTemp] = configsIni.FolderMigrations + file.Name()
				migrationPathTest := migrationPath{
					key:  strTemp,
					path: db.Migrates.migrations[strTemp],
				}
				db.Migrates.migrationsOrders = append(db.Migrates.migrationsOrders, migrationPathTest)
			}
		}
	}

	db.orderMigrationPath()
	fmt.Println("Folder Migrate:", configsIni.FolderMigrations)
	fmt.Printf("Migrates:\n\t%+v\n", db.Migrates.migrations)
}

func (db *DbInstance) initDB() {
	// Configiuracion de conecciones idle
	db.db().SetConnMaxLifetime(time.Duration(configsIni.SetConnMaxLifetime) * time.Minute)
	db.db().SetMaxIdleConns(configsIni.SetMaxIdleConns)
	db.db().SetMaxOpenConns(configsIni.SetMaxOpenConns)

	// Cargar de Migration
	//rand.Seed(time.Now().UnixNano())

}

// Ejecutar migration dado un nombre
func (db *DbInstance) execOneMigration(nombreMigration string) {
	file, err := os.Open(db.Migrates.migrations[nombreMigration])
	if err != nil {
		logg.GeneralLogger.Printf("fallo al abrir file %s relaccionado con migration:%s,  Error: %v \n", nombreMigration, db.Migrates.migrations[nombreMigration], nombreMigration, err)
		fmt.Printf("fallo al abrir file %s relaccionado con migration:%s,  Error: %v \n", nombreMigration, db.Migrates.migrations[nombreMigration], nombreMigration, err)
		return
	}
	defer file.Close()

	querieSql, err := io.ReadAll(file)
	if err != nil {
		logg.GeneralLogger.Printf("fallo al leer file relaccionado con :%s, Path: %s Error: %v \n", nombreMigration, db.Migrates.migrations[nombreMigration], err)
		fmt.Printf("fallo al leer file relaccionado con :%s, Path: %s Error: %v \n", nombreMigration, db.Migrates.migrations[nombreMigration], err)
		return
	}

	// Ejecuta las consultas SQL
	db.Exec(string(querieSql))
	if db.Error != nil {
		logg.GeneralLogger.Printf("Ocurrio un fallo ejecutando Query Error:%v \n", err)
		fmt.Printf("Ocurrio un fallo ejecutando Query Error:%v \n", err)
		return
	}
}

// Carga todas las migration situada en el directorio de migration
func (db *DbInstance) execAllMigration() {
	fmt.Println("Ejecutando todas las migartion:")
	logg.GeneralLogger.Println("Ejecutando todas las migartion:")
	fmt.Println("-----------------------------------------------")
	logg.GeneralLogger.Println("-----------------------------------------------")
	cantMig := 0
	for i := range db.Migrates.migrationsOrders {
		cantMig++
		db.execOneMigration(db.Migrates.migrationsOrders[i].key)
		fmt.Println("Migration:\033[36m", db.Migrates.migrationsOrders[i].key, "\033[0m")
		logg.GeneralLogger.Println("Migration:\033[36m", db.Migrates.migrationsOrders[i].key, "\033[0m")
	}
	fmt.Printf("------------End-------%s%d ----------------------------\n", "Cantida Migraciones ejecutadas en BD: ", cantMig)
	logg.GeneralLogger.Printf("------------End-------%s%d ----------------------------\n", "Cantida Migraciones ejecutadas en BD: ", cantMig)
}

func (db *DbInstance) execAllMigrationFromModels() {
	fmt.Println("\nEjecutando todas las migartion desde Models:")
	logg.GeneralLogger.Println("Ejecutando todas las migartion:")
	fmt.Println("-----------------------------------------------")
	logg.GeneralLogger.Println("-----------------------------------------------")
	cantMig := len(models.Models)
	db.AutoMigrate(models.Models...)
	// fmt.Println("Migration:\033[36m", reflect.TypeOf(v).String(), "\033[0m")
	// logg.GeneralLogger.Println("Migration:\033[36m", reflect.TypeOf(v).String(), "\033[0m")
	//}
	fmt.Printf("------------End-------%s%d ----------------------------\n", "Cantida Migraciones ejecutadas en BD: ", cantMig)
	logg.GeneralLogger.Printf("------------End-------%s%d ----------------------------\n", "Cantida Migraciones ejecutadas en BD: ", cantMig)
}

// Cargara grupo de migration especificadas que deben estar en el dir de migrations
func (db *DbInstance) ExecSetMigration(nombreMigration []string) {
	fmt.Println("Ejecutando todas las migartion:")
	logg.GeneralLogger.Println("Ejecutando todas las migartion:")
	fmt.Println("-----------------------------------------------")
	logg.GeneralLogger.Println("-----------------------------------------------")
	for i := range nombreMigration {
		db.execOneMigration(nombreMigration[i])
		fmt.Println("Migration:\033[36m", i, "\033[0m")
		logg.GeneralLogger.Println("Migration:\033[36m", i, "\033[0m")
	}
	fmt.Println("------------End-----------------------------------")
	logg.GeneralLogger.Println("------------End-----------------------------------")
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
