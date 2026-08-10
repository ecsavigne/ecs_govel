package partition

import (
	"ecs_govel/pkg/pkglog"

	"gorm.io/gorm"
)

func ExecutePartition(db *gorm.DB, partition ...string) {
	for _, sql := range partition {
		if res := db.Exec(sql); res.Error != nil {
			pkglog.Log.Errorf("Error executing partition: %v\n", res.Error)
		}
	}
}
