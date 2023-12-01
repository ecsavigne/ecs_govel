package repositories

import (
	"ecs_govel/app/models"
	"ecs_govel/config/db"
)

type EstudianteRepository struct {
}

// Metodos del Repositorio
func (e *EstudianteRepository) RegistrarEstudiante(persona *models.Persona, siJuridico bool) interface{} {
	estudiante := models.Estudiante{CI: persona.CI, SiJuridico: siJuridico}
	if err := db.Orm.Create(persona).Error; err != nil {
		panic("[Repositories.EstudianteRepository.RegistrarEstudiante: Line 15] - " + err.Error())
	} else if err := db.Orm.Create(&estudiante).Error; err != nil {
		panic("[Repositories.EstudianteRepository.RegistrarEstudiante: Line 17] - " + err.Error())
	}
	return true
}

func (e *EstudianteRepository) RegistrarEstudianteEnCurso(matricula *models.Matricula) interface{} {
	if err := db.Orm.Create(matricula).Error; err != nil {
		panic("[Repositories.EstudianteRepository.RegistrarEstudianteEnCurso: Line 24] - " + err.Error())
	}
	return true
}

func (e *EstudianteRepository) ModificarEstudiante(persona *models.Persona, siJuridico bool) interface{} {
	estudiante := models.Estudiante{
		CI:         persona.CI,
		SiJuridico: siJuridico,
	}
	if res := db.Orm.Model(new(models.Persona)).Updates(persona); res.Error != nil {
		panic("[Repositories.EstudianteRepository.ModificarEstudiante: Line 35] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	} else if res := db.Orm.Model(new(models.Estudiante)).Updates(&estudiante); res.Error != nil {
		panic("[Repositories.EstudianteRepository.ModificarEstudiante: Line 39] - " + res.Error.Error())
	}
	return true
}

func (e *EstudianteRepository) ModificarEstudianteCurso(matricula *models.Matricula) interface{} {
	if res := db.Orm.Model(new(models.Matricula)).Updates(matricula); res.Error != nil {
		panic("[Repositories.EstudianteRepository.ModificarEstudianteCurso: Line 46] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (e *EstudianteRepository) EliminarEstudiante(estudiante *models.Estudiante) interface{} {
	if res := db.Orm.Model(new(models.Estudiante)).Unscoped().Delete(estudiante); res.Error != nil {
		panic("[Repositories.EstudianteRepository.EliminarEstudiante: Line 55] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (e *EstudianteRepository) EliminarEstudianteCurso(matricula *models.Matricula) interface{} {
	if res := db.Orm.
		Unscoped().
		Where("estudiante_ci = ?", matricula.EstudianteCi).
		Where("curso_id = ?", matricula.CursoId).
		Delete(matricula); res.Error != nil {
		panic("[Repositories.EstudianteRepository.EliminarEstudianteCurso: Line 64] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (e *EstudianteRepository) MostrarEstudianteJuridicoNoJuridico(estudiante *models.Estudiante) interface{} {
	estudiantes := []map[string]interface{}{}
	if res := db.Orm.
		Table("estudiantes").
		Select("personas.*, estudiantes.si_juridico").
		Joins("inner join personas on personas.ci = estudiantes.ci").
		Where("si_juridico = ?", estudiante.SiJuridico).Scan(&estudiantes); res.Error != nil {
		panic("[Repositories.EstudianteRepository.mostrarEstudianteJuridicoNoJuridico: Line 74] - " + res.Error.Error())
	}
	return estudiantes
}

func (e *EstudianteRepository) MostrarEstudianteDeCurso(idCurso int) interface{} {
	estuds := []map[string]interface{}{}
	if res := db.Orm.
		Table("cursos").
		Select("cursos.id as idCurso, contenidos.tema, personas.*").
		Joins("inner join matriculas on matriculas.curso_id = cursos.id").
		Joins("inner join curso_contenidos on curso_contenidos.curso_id = cursos.id").
		Joins("inner join contenidos on curso_contenidos.contenido_id = contenidos.id").
		Joins("inner join estudiantes on matriculas.estudiante_ci = estudiantes.ci").
		Joins("inner join personas on estudiantes.ci = personas.ci").
		Order("personas.nombre asc").
		Find(&estuds, "cursos.id = ?", idCurso); res.Error != nil {
		panic("Ocurried one error in line 86 [EstudanteRepository.MostrarEstudianteDeCurso] error: " + res.Error.Error())
	}
	return estuds
}

func (e *EstudianteRepository) MostrarEstudiante() interface{} {
	pers := []models.Persona{}
	if res := db.Orm.
		Joins("inner join estudiantes on personas.ci = estudiantes.ci").
		Find(&pers); res.Error != nil {
		panic("Ocurried one error in line 86 [EstudanteRepository.mostrarEstudiante] error: " + res.Error.Error())
	}
	return pers
}
