package configs

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"ecs_govel/app/model"
	"ecs_govel/database/migration"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

func create_database_postgres() error {
	db, err := sql.Open("postgres", PG_DNS_DB)
	if err != nil {
		l := fmt.Sprintf("Error al conectar al servidor de PostgreSQL: %v", err)
		return fmt.Errorf("%s", l)
	}
	defer db.Close()

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = '%s')", PG_DB_NAME)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return fmt.Errorf("Error obtaining database existence status is: %s", err.Error())
	}

	if !exists {
		fmt.Println("Creating database")
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", PG_DB_NAME))
		if err != nil {
			return fmt.Errorf("Error creating database is: %s", err.Error())
		}
		fmt.Printf("Database '%s' created successfully.\n", PG_DB_NAME)
		// Grant access
		fmt.Println("Granting access to database")
		_, err = db.Exec(fmt.Sprintf("GRANT ALL ON DATABASE %s TO %s;", PG_DB_NAME, PG_DB_USER))
		if err != nil {
			return fmt.Errorf("Error granting database access is: %s", err.Error())
		}
		fmt.Printf("Access granted to database: '%s' by user: '%s'.\n", PG_DB_NAME, PG_DB_USER)
	} else {
		fmt.Printf("Database '%s' already exists.\n", PG_DB_NAME)
	}

	return nil
}

func logDBInfo() logger.Interface {
	logDataBaseFile, err := os.Create("database.log")
	if err != nil {
		fmt.Println("Error creating Database log file, is: ", err)
	}

	newLogger := logger.New(
		log.New(logDataBaseFile, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	return newLogger
}

// Init : First Setup
func (db *DbInstance) MigratePG() {
	logMessage := ""
	debugMessage := "1"
	defer func() {
		if r := recover(); r != nil {
			logMessage = filepath.Base(os.Args[0]) + " :  Error initializing Migration in Database. " + ": Recovered from exception " + ". Interface in defer is: " + fmt.Sprintf("%+v", r) + ". DebugMessage is: " + debugMessage
			Log.Debugf("[database.database.go - Init()]. ", logMessage)
			fmt.Println(logMessage)
		}
	}()

	debugMessage = "2"

	// db.DB.Migrator().DropTable(&migration.TestMigation{})
	err := db.DB.AutoMigrate(
		&migration.TestMigation{},
	)

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	// Load Migration from .sql
	// migration.ExecuteMigrationFromSql(db.DB, Log, true)

	// Load Seeders
	// seeders.ExecuteSeeders(db.DB)

	// execute trigger
	// trigger.ExecuteTrigger(db.DB, trigger.TriggerDeleteMessageByCompanyWhatsapp())

	// Execute Partition
	// partition.ExecutePartition(db.DB, partition.OpenConversationPartition())
}

func postgresDB() {
	var (
		logMessage, debugMessage string
		err                      error
	)

	defer func() {
		if r := recover(); r != nil {
			logMessage = filepath.Base(os.Args[0]) + " :  Error initializing Database. " + ": Recovered from exception " + ". Interface in defer is: " + fmt.Sprintf("%+v", r) + ". DebugMessage is: " + debugMessage
			Log.Errorf("[database.database.go - init()]. ", logMessage)
			//TODO: Quitar el Exit(2)
			os.Exit(2)
		}
	}()

	// Create database si no existe
	if err = create_database_postgres(); err != nil {
		logMessage = filepath.Base(os.Args[0]) + " :  Error initializing Database. " + err.Error()
		fmt.Println(logMessage)
		Log.Errorf("[database.database.go - init()]. ", logMessage)
		panic(logMessage)
	}

	PG_DB_CONNSTR = fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", PG_DB_HOST, PG_DB_USER, PG_DB_NAME, FORWARD_DB_PORT, PG_DB_PASSWORD)
	postgresSource := postgres.Open(PG_DB_CONNSTR)
	postgresReplica := postgres.Open(PG_DB_CONNSTR)

	Database.DB, err = gorm.Open(postgresSource, &gorm.Config{
		Logger: logDBInfo(),
	})

	// Create connection pool
	Database.Use(dbresolver.Register(dbresolver.Config{
		Sources:           []gorm.Dialector{postgresSource},
		Replicas:          []gorm.Dialector{postgresReplica},
		Policy:            dbresolver.RandomPolicy{},
		TraceResolverMode: true,
	} /*, &migration.Chat{}*/))

	dateFormat := time.Now()
	now := dateFormat.Format("2006-01-02 15:04:05")
	fmt.Println("New postgres conennection opened at ", now)
	if err != nil {
		panic(err)
	}
	model.SetGlobalDB(Database.DB)
	fmt.Printf("Base Datos >>>>>>>>>>>>>. : %+v\n", model.GetGlobalDB())

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
	// 	StartServer:     false,E
	// 	//MetricsCollector: []prometheus.MetricsCollector{},
	// }))

	// Load Migration
	Database.MigratePG()
}

func prepare_db() {
	switch DB_TYPE {
	case "postgres":
		postgresDB()
	case "mongo":
		mongoDB()
	case "all":
		mongoDB()
		postgresDB()
	default:
		postgresDB()
	}
}
