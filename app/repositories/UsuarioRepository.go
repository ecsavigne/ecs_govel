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
	//usr := new(models.Usuario)
	// config.Orm.Find(usr, uint(idUsuario))
	// usr.Password = newPass
	// if err := config.Orm.Save(usr).Error; err != nil {
	// 	return err
	// }
	return true
}

func (u UsuarioRepository) RecuperarPass(correio string) interface{} {
	usr := new(models.Usuario)
	usr.Password = "12345" // Aqui gneral un pass y mandarlo por correo
	// config.Orm.Model(usr).Where("mail = ?", correio).Save(*usr)
	return true
}
