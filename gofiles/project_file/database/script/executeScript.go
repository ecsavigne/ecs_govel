package script

import (
	"ecs_govel/rest/app/model"

	"gorm.io/gorm"
)

func ExecuteScript(db *gorm.DB, trigger ...string) {
	for _, sql := range trigger {
		if res := db.Exec(sql); res.Error != nil {
			model.Log.Errorf("Error executing Script: %v\n", res.Error)
		}
	}
}
