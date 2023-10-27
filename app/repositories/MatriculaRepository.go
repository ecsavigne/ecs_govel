package repositories

import (
	"ecs_govel/app/models"
	"ecs_govel/config/db"
)

type MatriculaRepository struct {
}

// Metodos del Repositorio
// se le pasa una matricula con id de curso, muestra matriculas de un curso
func (m *MatriculaRepository) MostrarMatricula(matricula *models.Matricula) interface{} {
	matriculas := make([]models.Matricula, 0)
	if res := db.Orm.Model(new(models.Matricula)).Find(&matriculas, matricula); res.Error != nil {
		panic("[Repositories.MatriculaRepository.MostrarMatricula : Line 17] - " + res.Error.Error())
	}
	return matriculas
}
