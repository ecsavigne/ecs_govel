package seeders

import (
	"ecs_govel/app/model"

	"gorm.io/gorm"
)

type Seeders struct {
	*gorm.DB
}

func NewSeeders(dBase ...*gorm.DB) *Seeders {
	db := &gorm.DB{}
	if len(dBase) != 0 {
		db = dBase[0]
	}
	s := new(Seeders)
	s.DB = db
	return s
}

func (s *Seeders) run() {
	if s.DB == nil {
		model.Log.Errorf("Not must run seeders, DB is nil.\n")
		return
	}
	// s.applicationSeeder()
	// s.companyWhatsappSeeder()
	s.TestSeeder()
}

func ExecuteSeeders(db *gorm.DB) {
	seed := NewSeeders(db)
	seed.run()
}
