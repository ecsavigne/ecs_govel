package script

import (
	"ecs_govel/pkg/pkglog"

	"gorm.io/gorm"
)

func ExecuteScript(db *gorm.DB, trigger ...string) {
	for _, sql := range trigger {
		if res := db.Exec(sql); res.Error != nil {
			pkglog.Log.Errorf("Error executing Script: %v\n", res.Error)
		}
	}
}
