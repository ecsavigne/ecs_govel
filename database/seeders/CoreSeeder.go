package seeders

import (
	"gorm.io/gorm"
)

type Seeders struct {
	*gorm.DB
}

func NewSeeders(db *gorm.DB) *Seeders {
	s := new(Seeders)
	s.DB = db
	return s
}

func (s *Seeders) Run() {
	s.TestSeeder()
}
