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
	"ecs_govel/database/partition"
	"ecs_govel/database/seeder"
	"ecs_govel/database/shared"
	"ecs_govel/pkg/pkglog"

	logecs "github.com/ecsavigne/logecs/log"
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
		pkglog.Log.Create(logecs.InfoLog{Type: logecs.Error, Sub: "drivergorm.connectDBPostgres", Name: "connect_db_open", Content: map[string]any{"date_time": time.Now().Format("2006-01-02 15:04:05"), "type_db": c_.DB_TYPE, "error": fmt.Sprintf("%s%s%s", "\033[31m", err.Error(), "\033[0m")}})
	}

	dialectorPostgress := postgres.New(postgres.Config{
		Conn: sqlDB,
	})
	db, err := gorm.Open(dialectorPostgress, &gorm.Config{})
	if err != nil {
		pkglog.Log.Create(logecs.InfoLog{Type: logecs.Error, Sub: "drivergorm.connectDBPostgres", Name: "connect_db_postgres", Content: map[string]any{"date_time": time.Now().Format("2006-01-02 15:04:05"), "type_db": c_.DB_TYPE, "error": fmt.Sprintf("%s%s%s", "\033[31m", err.Error(), "\033[0m")}})
	} else {
		pkglog.Log.Create(logecs.InfoLog{Type: logecs.Info, Sub: "drivergorm.connectDBPostgres", Name: "connect_db_postgres", Content: map[string]any{"date_time": time.Now().Format("2006-01-02 15:04:05"), "type_db": c_.DB_TYPE}})
	}

	dateFormat := time.Now()
	now := dateFormat.Format("2006-01-02 15:04:05")

	pkglog.Log.Create(logecs.InfoLog{Type: logecs.Info, Sub: "drivergorm.connectDBPostgres", Name: "connect_db_postgres_exit", Content: map[string]any{"date_time": now, "type_db": c_.DB_TYPE, "name_connect": nameConnect}})

	return db
}

func logDBInfo() logger.Interface {
	logDataBaseFile, err := os.Create(filepath.Join(c_.PATH_BASE, "database.log"))
	if err != nil {
		pkglog.Log.Create(logecs.InfoLog{Type: logecs.Error, Sub: "drivergorm.logDBInfo", Name: "log_db_info", Content: map[string]any{"date_time": time.Now().Format("2006-01-02 15:04:05"), "type_db": c_.DB_TYPE, "error": err.Error()}})
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
	debugMessage := "1"
	defer func() {
		if r := recover(); r != nil {
			pkglog.Log.Create(logecs.InfoLog{Type: logecs.Error, Sub: "drivergorm.migratePG", Name: "migrate_pg__recovery_error", Content: map[string]any{"date_time": time.Now().Format("2006-01-02 15:04:05"), "type_db": c_.DB_TYPE, "error": fmt.Sprintf("%+v", r), "debug_message": debugMessage}})
		}
	}()

	debugMessage = "2"

	// dbManager.DB.Migrator().DropTable(&migration.TestMigation{})
	// dbManager.DB.Migrator().DropTable(&migration.TestMigation{})
	// Execute Script
	partition.New(dbManager.DB).Run()

	err := dbManager.DB.AutoMigrate(
	// &migration.AllowCompany{},
	// &migration.Company{},
	// &migration.IGAccount{},
	// new(migration.Message),
	)

	if err != nil {
		pkglog.Log.Create(logecs.InfoLog{Type: logecs.Error, Sub: "drivergorm.migratePG", Name: "migrate_pg_auto_migrate", Content: map[string]any{"date_time": time.Now().Format("2006-01-02 15:04:05"), "type_db": c_.DB_TYPE, "error": err.Error()}})
	} else {
		pkglog.Log.Create(logecs.InfoLog{Type: logecs.Info, Sub: "drivergorm.migratePG", Name: "migrate_pg_auto_migrate", Content: map[string]any{"date_time": time.Now().Format("2006-01-02 15:04:05"), "type_db": c_.DB_TYPE, "msg": "Load migration successfully"}})
	}

	// Load Migration from .sql
	// migration.ExecuteMigrationFromSql(dbManager.DB, Log, true)

	// Load Seeders
	seeder.ExecuteSeeders(dbManager.DB)

	// execute trigger
	// trigger.ExecuteTrigger(dbManager.DB, trigger.TriggerDeleteMessageByCompanyWhatsapp())

	// partition.ExecutePartition(dbManager.DB, partition.OpenConversationPartition())
}

// db postgres principal
func PostgresDB(managerGormDB *shared.DBManager) {
	var (
		debugMessage string = "1"
		err          error
	)

	defer func() {
		if r := recover(); r != nil {
			pkglog.Log.Create(logecs.InfoLog{Type: logecs.Error, Sub: "drivergorm.connectDBPostgres", Name: "connect_db_open", Content: map[string]any{"date_time": time.Now().Format("2006-01-02 15:04:05"), "type_db": c_.DB_TYPE, "error": fmt.Sprintf("%+v", r), "debug_message": debugMessage}})
		}
	}()

	// Create database si no existe
	if err = create_database_postgres(); err != nil {
		debugMessage = "2"
		panic(err.Error())
	}

	debugMessage = "3"
	PG_DB_CONNSTR := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", c_.PG_DB_HOST, c_.PG_DB_USER, c_.PG_DB_NAME, c_.FORWARD_DB_PORT, c_.PG_DB_PASSWORD)
	postgresSource := postgres.Open(PG_DB_CONNSTR)
	postgresReplica := postgres.Open(PG_DB_CONNSTR)

	managerGormDB.DB, err = gorm.Open(postgresSource, &gorm.Config{
		Logger: logDBInfo(),
	})
	if err != nil {
		panic(err)
	}

	debugMessage = "4"
	// Create connection pool
	err = managerGormDB.Use(dbresolver.Register(dbresolver.Config{
		Sources:           []gorm.Dialector{postgresSource},
		Replicas:          []gorm.Dialector{postgresReplica},
		Policy:            dbresolver.RandomPolicy{},
		TraceResolverMode: true,
	} /*, &migration.Chat{}*/))
	if err != nil {
		panic(err)
	}

	debugMessage = "5"
	dateFormat := time.Now()
	now := dateFormat.Format("2006-01-02 15:04:05")
	pkglog.Log.Warnf("New postgres conennection opened at: %s\n", now)

	sqlDB, _ := managerGormDB.DB.DB()
	debugMessage = "6"
	sqlDB.SetConnMaxLifetime(time.Minute * 2) // Make than last forever
	sqlDB.SetMaxIdleConns((c_.APP_MAX_CONNECTIONS / c_.APP_CANT_X) - 7)
	sqlDB.SetMaxOpenConns((c_.APP_MAX_CONNECTIONS / c_.APP_CANT_X) - 5)
	debugMessage = "7"
	pkglog.Log.Infof("Max Connections: %d, CantX: %d\n", c_.APP_MAX_CONNECTIONS, c_.APP_CANT_X)

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
	debugMessage = "8"
	migratePG(managerGormDB)
}
