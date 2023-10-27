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
		panic("[Repositories.Usuarirepository.RegistrarUsr: Line 18] - " + err.Error())
	}
	return true
}

func (u UsuarioRepository) CambiarPass(idUsuario int, newPass string) interface{} {
	if res := db.Orm.Model(new(models.Usuario)).Where("id = ?", idUsuario).Update("password", newPass); res.Error != nil {
		panic("[Repositories.Usuarirepository.CambiarPass: Line 26] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (u UsuarioRepository) RecuperarPass(correio string) interface{} {
	//usr.Password = "12345" // Aqui gneral un pass y mandarlo por correo
	res := db.Orm.Model(new(models.Usuario)).Where("mail = ?", correio).Update("password", "12345")
	if res.Error != nil {
		panic("[Repositories.Usuarirepository.RecuperarPass: Line 37] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}
