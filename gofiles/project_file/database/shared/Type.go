package shared

import "gorm.io/gorm"

type DBManager struct {
	*gorm.DB
	AppID string
}
