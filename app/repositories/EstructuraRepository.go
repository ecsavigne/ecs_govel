package repositories

import (
	"ecs_govel/app/models"
	"ecs_govel/config/db"
)

type EstructuraRepository struct {
}

// Metodos del Repositorio
func (e *EstructuraRepository) RegistrarEstructura(estructura *models.Estructura) interface{} {
	if err := db.Orm.Create(estructura).Error; err != nil {
		panic("[Repositories.EstructuraRepository.RegistrarEstructura: Line 14] - " + err.Error())
	}
	return true
}

func (e *EstructuraRepository) MostrarEstructura(estructura *models.Estructura) interface{} {
	estructuras := models.Estructura{}
	if res := db.Orm.Model(new(models.Estructura)).Find(&estructuras, estructura); res.Error != nil {
		panic("[Repositories.EstructuraRepository.MostrarEstructura: Line 22] - " + res.Error.Error())
	}
	return estructuras
}

func (e *EstructuraRepository) EliminarEstructura(estructura *models.Estructura) interface{} {
	if res := db.Orm.Model(new(models.Estructura)).Unscoped().Delete(estructura); res.Error != nil {
		panic("[Repositories.EstructuraRepository.EliminarEstructura: Line 29] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (e *EstructuraRepository) ModificarEstructura(estructura *models.Estructura) interface{} {
	if res := db.Orm.Model(new(models.Estructura)).Updates(estructura); res.Error != nil {
		panic("[Repositories.EstructuraRepository.ModificarEstructura: Line 38] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (e *EstructuraRepository) MostrarAllEstructura() interface{} {
	estructuras := models.Estructura{}
	if res := db.Orm.Model(new(models.Estructura)).Find(&estructuras); res.Error != nil {
		panic("[Repositories.EstructuraRepository.MostrarEstructura: Line 48] - " + res.Error.Error())
	}
	return estructuras
}
