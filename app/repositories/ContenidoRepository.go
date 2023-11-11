package repositories

import (
	"ecs_govel/app/models"
	"ecs_govel/config/db"
	"fmt"
)

type ContenidoRepository struct {
}

// Metodos del Repositorio
// Registrar Contenido resive un objeto que tiene todos los datos que se necesistan
// Para el registro
func (c *ContenidoRepository) RegistrarContenido(contenido *models.Contenido) interface{} {
	if err := db.Orm.Create(contenido).Error; err != nil {
		panic("[Repositories.ContenidoRepository.RegistrarContenido: Line 16] - " + err.Error())
	}
	return true
}

func (c *ContenidoRepository) AsociarContenidoToEstructura(estructura_contenido *models.EstructuraContenido) interface{} {
	if err := db.Orm.Create(estructura_contenido).Error; err != nil {
		panic("[Repositories.ContenidoRepository.AsociarContenidoToEstructura: Line 23] - " + err.Error())
	}
	return true
}

func (c *ContenidoRepository) ModificarContenidoToEstructura(estructura_contenido *models.EstructuraContenido) interface{} {
	if res := db.Orm.Model(new(models.EstructuraContenido)).Updates(estructura_contenido); res.Error != nil {
		panic("[Repositories.ContenidoRepository.ModificarContenidoToEstructura: Line 30] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (c *ContenidoRepository) MostrarContenido(contenido *models.Contenido) interface{} {
	contenidos := models.Contenido{}
	if res := db.Orm.Model(&contenidos).Find(&contenidos, contenido); res.Error != nil {
		panic("[Repositories.ContenidoRepository.MostrarContenido: Line 40] - " + res.Error.Error())
	}
	return contenidos
}

func (c *ContenidoRepository) MostrarAllContenido() interface{} {
	contenidos := make([]models.Contenido, 0)
	if res := db.Orm.Model(new(models.Contenido)).Find(&contenidos); res.Error != nil {
		panic("[Repositories.ContenidoRepository.MostrarAllContenido: Line 48] - " + res.Error.Error())
	}
	fmt.Printf("Conetnidos: \n%+v\n", contenidos)
	return contenidos
}

func (c *ContenidoRepository) EliminarContenido(contenido *models.Contenido) interface{} {
	if res := db.Orm.Model(new(models.Contenido)).Delete(contenido); res.Error != nil {
		panic("[Repositories.ContenidoRepository.EliminarContenido: Line 55] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (c *ContenidoRepository) ModificarContenido(contenido *models.Contenido) interface{} {
	if res := db.Orm.Model(new(models.Contenido)).Updates(contenido); res.Error != nil {
		panic("[Repositories.ContenidoRepository.ModificarContenido: Line 64] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}
