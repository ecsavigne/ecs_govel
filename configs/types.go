package configs

import "gorm.io/gorm"

type TYPE_LOG_MESSAGE = string

type DbInstance struct {
	*gorm.DB
	AppID string
}
