package repositories

import (
	"ecs_govel/app/models"
	"ecs_govel/config/db"
)

type CursoRepository struct {
}

// Metodos del Repositorio
func (c *CursoRepository) RegistrarCurso(curso *models.Curso, curso_contenido *models.CursoContenido) interface{} {
	if res := db.Orm.Create(curso); res.Error != nil {
		panic("[Repositories.CursoRepository.RegistrarCurso: Line 14] - " + res.Error.Error())
	} else {
		curso_contenido.CursoId = curso.ID
		if err := db.Orm.Create(curso_contenido).Error; err != nil {
			panic("[Repositories.CursoRepository.RegistrarCurso: Line 18] - " + err.Error())
		}
	}
	return true
}

func (c *CursoRepository) ModificarCurso(curso *models.Curso, curso_contenido *models.CursoContenido) interface{} {
	if res := db.Orm.Model(new(models.Curso)).Updates(curso); res.Error != nil {
		panic("[Repositories.CursoRepository.RegistrarCurso: Line 26] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (c *CursoRepository) EliminarCurso(curso *models.Curso) interface{} {
	if res := db.Orm.Model(new(models.Curso)).Delete(curso); res.Error != nil {
		panic("[Repositories.CursoRepository.EliminarCurso: Line 35] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (c *CursoRepository) MostrarCurso(curso *models.Curso) interface{} {
	cursos := make([]models.Curso, 0)
	if res := db.Orm.Model(new(models.Curso)).Find(&cursos, curso); res.Error != nil {
		panic("[Repositories.CursoRepository.MostrarCurso: Line 45] - " + res.Error.Error())
	}
	return cursos
}

func (c *CursoRepository) MostrarCursos() interface{} {
	curs := []models.Curso{}
	if res := db.Orm.Find(&curs); res.Error != nil {
		panic("Ocurried one error in line 53 [CursoRepository.MostrarCursos] error: " + res.Error.Error())
	}
	return curs
}
