package seeders

import (
	"context"
	"ecs_govel/app/models"
)

func (s *Seeders) SeederUsuario() {
	usr := &models.Usuario{
		ID:       14,
		Nombre:   "julio",
		Password: "3333",
		Mail:     "ecsavigne@gmail.com",
	}

	if res := s.Save(usr); res.Error != nil {
		Seeder.Logger.Error(context.Background(), "Error: "+res.Error.Error())
	}
	Seeder.Logger.Info(context.Background(), "Seeder User ejecutrado satifactoriamente")
}
