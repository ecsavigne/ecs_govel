package drivergorm

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	c_ "ecs_govel/configs"
	"ecs_govel/database/script"
	"ecs_govel/database/shared"
	"ecs_govel/pkg/pkglog"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

func create_database_postgres() error {
	db, err := sql.Open("postgres", c_.PG_DNS_DB)
	if err != nil {
		l := fmt.Sprintf("Error al conectar al servidor de PostgreSQL: %v", err)
		return fmt.Errorf("%s", l)
	}
	defer db.Close()

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = '%s')", c_.PG_DB_NAME)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return fmt.Errorf("Error obtaining database existence status is: %s", err.Error())
	}

	if !exists {
		fmt.Println("Creating database")
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", c_.PG_DB_NAME))
		if err != nil {
			return fmt.Errorf("Error creating database is: %s", err.Error())
		}
		fmt.Printf("Database '%s' created successfully.\n", c_.PG_DB_NAME)
		// Grant access
		fmt.Println("Granting access to database")
		_, err = db.Exec(fmt.Sprintf("GRANT ALL ON DATABASE %s TO %s;", c_.PG_DB_NAME, c_.PG_DB_USER))
		if err != nil {
			return fmt.Errorf("Error granting database access is: %s", err.Error())
		}
		fmt.Printf("Access granted to database: '%s' by user: '%s'.\n", c_.PG_DB_NAME, c_.PG_DB_USER)
	} else {
		fmt.Printf("Database '%s' already exists.\n", c_.PG_DB_NAME)
	}

	return nil
}

// create connection postgres, utils for create connection to db postgres not principal
func connectDBPostgres(nameConnect, dns string) *gorm.DB {
	sqlDB, err := sql.Open("pgx", dns)
	if err != nil {
		pkglog.Log.Sub("Db").Infof("\033[31mError connecting to %s using sql package, error is: %s\033[0m\n", nameConnect, err.Error())
	}

	dialectorPostgress := postgres.New(postgres.Config{
		Conn: sqlDB,
	})
	db, err := gorm.Open(dialectorPostgress, &gorm.Config{})
	if err != nil {
		pkglog.Log.Sub("Db").Errorf("\033[31mError connecting to %s: error is: %s\033[0m\n", nameConnect, err.Error())
	} else {
		pkglog.Log.Sub("Db").Infof("\033[34mLoaded " + nameConnect + " db\033[0m\n")
	}

	dateFormat := time.Now()
	now := dateFormat.Format("2006-01-02 15:04:05")
	pkglog.Log.Sub("Db").Infof("New postgres conennection %s opened at %s\n", nameConnect, now)

	return db
}

func logDBInfo() logger.Interface {
	logDataBaseFile, err := os.Create(filepath.Join(c_.PATH_BASE, "database.log"))
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

// for run migration, script, seeder in db postgres principal
func migratePG(dbManager *shared.DBManager) {
	logMessage := ""
	debugMessage := "1"
	defer func() {
		if r := recover(); r != nil {
			logMessage = filepath.Base(os.Args[0]) + " :  Error initializing Migration in managerGormDB. " + ": Recovered from exception " + ". Interface in defer is: " + fmt.Sprintf("%+v", r) + ". DebugMessage is: " + debugMessage
			pkglog.Log.Debugf("[database.database.go - Init()]. ", logMessage)
			fmt.Println(logMessage)
		}
	}()

	debugMessage = "2"

	// dbManager.DB.Migrator().DropTable(&migration.TestMigation{})
	// dbManager.DB.Migrator().DropTable(&migration.TestMigation{})
	// Execute Script
	script.ExecuteScript(dbManager.DB,
		script.CreateConversationsPartitionYear(),
	)
	err := dbManager.DB.AutoMigrate(
	// &migration.AllowCompany{},
	// &migration.Company{},
	// &migration.IGAccount{},
	// new(migration.Message),
	)

	if err != nil {
		pkglog.Log.Errorf("Error in migration: %s\n", err.Error())
	} else {
		pkglog.Log.Infof("Load migration successfully. \n")
	}

	// Load Migration from .sql
	// migration.ExecuteMigrationFromSql(dbManager.DB, Log, true)

	// Load Seeders
	// seeders.ExecuteSeeders(dbManager.DB)

	// execute trigger
	// trigger.ExecuteTrigger(dbManager.DB, trigger.TriggerDeleteMessageByCompanyWhatsapp())

	// partition.ExecutePartition(dbManager.DB, partition.OpenConversationPartition())
}

// db postgres principal
func PostgresDB(managerGormDB *shared.DBManager) {
	var (
		logMessage, debugMessage string
		err                      error
	)

	defer func() {
		if r := recover(); r != nil {
			logMessage = filepath.Base(os.Args[0]) + " :  Error initializing managerGormDB. " + ": Recovered from exception " + ". Interface in defer is: " + fmt.Sprintf("%+v", r) + ". DebugMessage is: " + debugMessage
			pkglog.Log.Sub("DbInstance").Debugf(`MigratePG {panic: "%s" }. %s`, logMessage, "\n")
		}
	}()

	// Create database si no existe
	if err = create_database_postgres(); err != nil {
		logMessage = filepath.Base(os.Args[0]) + " :  Error initializing managerGormDB. " + err.Error()
		fmt.Println(logMessage)
		pkglog.Log.Errorf("[database.database.go - init()]. %s\n", logMessage)
		panic(logMessage)
	}

	PG_DB_CONNSTR := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", c_.PG_DB_HOST, c_.PG_DB_USER, c_.PG_DB_NAME, c_.FORWARD_DB_PORT, c_.PG_DB_PASSWORD)
	postgresSource := postgres.Open(PG_DB_CONNSTR)
	postgresReplica := postgres.Open(PG_DB_CONNSTR)

	managerGormDB.DB, err = gorm.Open(postgresSource, &gorm.Config{
		Logger: logDBInfo(),
	})

	// Create connection pool
	managerGormDB.Use(dbresolver.Register(dbresolver.Config{
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

	pkglog.Log.Debugf("Max Connections: ", c_.APP_MAX_CONNECTIONS, " CantX: ", c_.APP_CANT_X)
	sqlDB, _ := managerGormDB.DB.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 2) // Make than last forever
	sqlDB.SetMaxIdleConns((c_.APP_MAX_CONNECTIONS / c_.APP_CANT_X) - 7)
	sqlDB.SetMaxOpenConns((c_.APP_MAX_CONNECTIONS / c_.APP_CANT_X) - 5)

	rand.New(rand.NewSource(time.Now().UnixNano()))
	// Agregar plugin para prometeus
	//fmt.Println("aAse datos mericas:", fmt.Sprintf("%s:%s", HTTP_SERVER_HOST_METRICS, HTTP_SERVER_PORT_METRICS))
	// managerGormDB.DB.Use(prometheus.New(prometheus.Config{
	// 	DBName:          DB_NAME,
	// 	RefreshInterval: 15,
	// 	PushAddr:        fmt.Sprintf("%s:%s", HTTP_SERVER_HOST_METRICS, HTTP_SERVER_PORT_METRICS),
	// 	StartServer:     false,E
	// 	//MetricsCollector: []prometheus.MetricsCollector{},
	// }))

	// Load Migration
	migratePG(managerGormDB)
}
