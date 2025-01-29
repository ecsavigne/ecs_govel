package model

import "ecs_govel/database/migration"

// User  : User model
type User migration.User

// type User struct {
// 	gorm.Model
// 	//TODO UserID ya se define por Defecto cuando se establece alguna Relaccion y el ID es feredado do Modelo
// 	//UserID   uint
// 	Username string `gorm:"not null;unique"`
// 	APIHash  string `gorm:"unique"`
// }
