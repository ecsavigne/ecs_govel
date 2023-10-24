package repositories

import (
	"ecs_govel/app/models"
	"ecs_govel/config/db"
)

type UsuarioRepository struct {
}

// Metodos del Repositorio
func (u UsuarioRepository) RegistrarUsr(nome, correio, pass string) interface{} {
	usr := new(models.Usuario)
	usr.Nombre = nome
	usr.Password = pass
	usr.Mail = correio
	if err := db.Orm.Create(usr).Error; err != nil {
		return err
	}
	return true
}

func (u UsuarioRepository) CambiarPass(idUsuario int, newPass string) interface{} {
	usr := new(models.Usuario)
	db.Orm.Find(usr, uint(idUsuario))
	usr.Password = newPass
	if err := db.Orm.Save(usr).Error; err != nil {
		return err
	}
	return true
}

func (u UsuarioRepository) RecuperarPass(correio string) interface{} {
	result := db.Orm.Model(new(models.Usuario)).Where("mail = ?", correio).Update("password", "12345")
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return false
	}
	return true
}
