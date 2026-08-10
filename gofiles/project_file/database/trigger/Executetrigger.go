package trigger

import (
	"ecs_govel/pkg/pkglog"

	"gorm.io/gorm"
)

func ExecuteTrigger(db *gorm.DB, trigger ...string) {
	for _, sql := range trigger {
		if res := db.Exec(sql); res.Error != nil {
			pkglog.Log.Errorf("Error executing trigger: %v\n", res.Error)
		}
	}
}
