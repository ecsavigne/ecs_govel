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

func (c *CursoRepository) ModificarCurso(curso *models.Curso, curso_contenido *models.CursoContenido, idContAsoc int) interface{} {
	if res := db.Orm.Model(new(models.Curso)).
		Where("id = ?", curso.ID).
		Updates(curso); res.Error != nil {
		panic("[Repositories.CursoRepository.RegistrarCurso: Line 27] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	} else {
		if res := db.Orm.Table("curso_contenidos").
			Where("curso_id = ?", curso_contenido.CursoId).
			Where("contenido_id = ?", idContAsoc).
			Updates(curso_contenido); res.Error != nil {
			panic("[Repositories.CursoRepository.RegistrarCurso: Line 35] - " + res.Error.Error())
		} else if res.RowsAffected == 0 {
			return false
		}
	}
	return true
}

func (c *CursoRepository) EliminarCurso(curso *models.Curso) interface{} {
	if res := db.Orm.Unscoped().Delete(curso); res.Error != nil {
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

func (c *CursoRepository) MostrarCursosContenido() interface{} {
	curs := []map[string]interface{}{}
	if res := db.Orm.
		Table("cursos").
		Select("cursos.id as idCurso, DATE_FORMAT(cursos.fecha_creacion, '%d-%m-%Y') AS fecha_creacion," +
			"cursos.duracion_hora,cursos.si_certificado," +
			"contenidos.id as idContenido, contenidos.tema").
		Joins("inner join curso_contenidos on cursos.id = curso_contenidos.curso_id").
		Joins("inner join contenidos on contenidos.id = curso_contenidos.contenido_id").
		Order("fecha_creacion asc, contenidos.tema asc").
		Scan(&curs); res.Error != nil {
		panic("Ocurried one error in line 63 [CursoRepository.MostrarCursosContenido] error: " + res.Error.Error())
	}
	return curs
}
