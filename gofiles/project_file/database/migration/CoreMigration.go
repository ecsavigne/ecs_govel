/*
1. Create migrate in path ./database/migration
2. Load file .sql and execute. Execute function: ExecuteMigrationFromSql(db *gorm.DB, log logecs.Logger, ifViewImportMigrations ...bool)
*/
package migration

import (
	"ecs_govel/pkg/pkglog"
	"fmt"
	"os"
	"path"
	"reflect"

	logecs "github.com/ecsavigne/logecs/log"

	"gorm.io/gorm"
)

type coreMigration struct {
	db                     *gorm.DB
	ifViewImportMigrations bool
	exec                   func(os.DirEntry)
}

// create new coreMigration
func New(db *gorm.DB) *coreMigration {
	return &coreMigration{db: db}
}

// Run execute all migrations
func (self coreMigration) Run() error {
	typ := reflect.TypeFor[*coreMigration]() // information of struct
	value := reflect.ValueOf(&self)          // value of struct

	for method := range typ.Methods() {
		methodValue := value.MethodByName(method.Name)

		if method.Name == "Migrate" {
			continue
		}

		if method.Func.IsValid() {
			pkglog.Log.Infof("executing migrator Method: %+v\n", method.Name)
			sliceValue := methodValue.Call([]reflect.Value{})

			if len(sliceValue) > 0 {
				v := sliceValue[0].Interface()
				if v != nil {
					if err, ok := v.(error); ok {
						pkglog.Log.Errorf("executed migrator Method: %v\n", err.Error())
					}
				}
			}
			pkglog.Log.Infof("executed migrator Method: %+v\n", method.Name)
		}
	}

	return nil
}

func (c coreMigration) migrateSQL() {
	str, _ := os.Getwd()
	path_migrations := fmt.Sprintf("%s/database/migration/", str)
	pkglog.Log.Infof("\x1b[33mLoad migrations from path:\x1b[0m \x1b[34m%s\x1b[0m\n", path_migrations)

	c.exec = func(f os.DirEntry) {
		content, e := os.ReadFile(fmt.Sprintf("%s%s", path_migrations, f.Name()))
		if e != nil {
			pkglog.Log.Errorf("Error read file migrations error is: \x1b[31m%s\x1b[0m\n", e.Error())
		}

		// c.Debugf("sql: \x1b[32m\n%s\x1b[0m\n\n", string(content))
		if err := c.db.Exec(string(content)).Error; err != nil {
			pkglog.Log.Errorf("Error executing migration error is: %s\n", err.Error())
		}
	}

	//  get all migrations from file.sql
	dirEntry, e := os.ReadDir(path_migrations)
	if e != nil {
		pkglog.Log.Errorf("Error read dir migrations error is: \x1b[31m%s\x1b[0m\n", e.Error())
	}

	for _, dir := range dirEntry {
		if dir.IsDir() || path.Ext(dir.Name()) != ".sql" {
			continue
		}

		if c.ifViewImportMigrations {
			pkglog.Log.Debugf("Found migration: \x1b[34m%s\x1b[0m\n", dir.Name())
		}

		c.exec(dir)
	}
}

// ExecuteMigrationFromSql executes all migrations from path ./database/migration/*.sql
// Log is a logger interface. If nil, no log will be printed.
// db is a database connection.
// Optional parameter ifViewImportMigrations if true, print all migrations found in the path.
// If false, do not print anything.
func ExecuteMigrationFromSql(db *gorm.DB, log logecs.Logger, ifViewImportMigrations ...bool) {
	view := false
	if len(ifViewImportMigrations) > 0 {
		view = ifViewImportMigrations[0]
	}
	core := coreMigration{db, view, nil}

	core.migrateSQL()
}
