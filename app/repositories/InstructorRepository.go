package repositories

import (
	"ecs_govel/app/models"
	"ecs_govel/config/db"
)

type InstructorRepository struct {
}

// Metodos del Repositorio
func (i *InstructorRepository) RegistrarInstructor(persona *models.Persona) interface{} {
	instructor := models.Instructore{CI: persona.CI}
	if err := db.Orm.Create(persona).Error; err != nil {
		panic("[Repositories.InstructorRepository.RegistrarInstructor: Line 15] - " + err.Error())
	} else if err := db.Orm.Create(&instructor).Error; err != nil {
		panic("[Repositories.InstructorRepository.RegistrarInstructor: Line 17] - " + err.Error())
	}
	return true
}

func (i *InstructorRepository) ModificarInstructor(instructor *models.Persona) interface{} {
	if res := db.Orm.Model(&instructor).Updates(instructor); res.Error != nil {
		panic("[Repositories.InstructorRepository.ModificarInstructor: Line 24] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (i *InstructorRepository) EliminarInstructor(instructor *models.Persona) interface{} {
	if res := db.Orm.Model(&instructor).Unscoped().Delete(instructor); res.Error != nil {
		panic("[Repositories.InstructorRepository.EliminarInstructor: Line 33] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

// Muestra datos personales de un instructor
func (i *InstructorRepository) MostrarInstructor(instructor *models.Instructore) interface{} {
	inst := models.Persona{}
	if res := db.Orm.Model(new(models.Persona)).Find(&inst, instructor); res.Error != nil {
		panic("[Repositories.InstructorRepository.MostrarInstructor: Line 44] - " + res.Error.Error())
	}
	return inst
}

func (i *InstructorRepository) MostrarInstructores() interface{} {
	inst := []map[string]interface{}{}
	if res := db.Orm.Table("instructores").
		Select("personas.*").
		Joins("inner join personas on personas.ci  = instructores.ci ").Find(&inst); res.Error != nil {
		panic("Ocurried one error in line 54 [EstudanteRepository.mostrarInstructores] error: " + res.Error.Error())
	}
	return inst
}

func (i *InstructorRepository) AgregarInstructorToCurso(instructor_curso *models.InstructorCurso) interface{} {
	if err := db.Orm.Create(instructor_curso).Error; err != nil {
		panic("[Repositories.InstructorRepository.AgregarInstructorToCurso: Line 59] - " + err.Error())
	}
	return true
}

func (i *InstructorRepository) EliminarCursoOffInstructor(instructor_curso *models.InstructorCurso) interface{} {
	if res := db.Orm.Model(new(models.InstructorCurso)).
		Unscoped().
		Where("instructor_ci = ?", instructor_curso.InstructorCi).
		Where("curso_id = ? ", instructor_curso.CursoId).
		Delete(instructor_curso); res.Error != nil {
		panic("[Repositories.InstructorRepository.EliminarInstructorDeCurso: Line 66] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (i *InstructorRepository) ModificarInstructorCurso(instructor_curso *models.InstructorCurso, idCursoA int) interface{} {
	if res := db.Orm.Model(new(models.InstructorCurso)).
		Where("instructor_ci = ?", instructor_curso.InstructorCi).
		Where("curso_id = ?", idCursoA).
		Where("curso_id <> ?", instructor_curso.CursoId).
		Updates(instructor_curso); res.Error != nil {
		panic("[Repositories.InstructorRepository.ModificarInstructorCurso: Line 75] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (i *InstructorRepository) MostrarCursosOfInstructor(ci string) interface{} {
	inst := []map[string]interface{}{}
	if res := db.Orm.Table("instructores").
		Select("cursos.id as idCurso, cursos.fecha_creacion, contenidos.tema").
		Joins("inner join personas on personas.ci  = instructores.ci ").
		Joins("inner join instructor_cursos on instructor_cursos.instructor_ci  = instructores.ci ").
		Joins("inner join cursos on instructor_cursos.curso_id  = cursos.id ").
		Joins("inner join curso_contenidos on curso_contenidos.curso_id  = cursos.id ").
		Joins("inner join contenidos on curso_contenidos.contenido_id  = contenidos.id ").
		Find(&inst, "personas.ci = ?", ci); res.Error != nil {
		panic("Ocurried one error in line 92 [InstructorRepository.MostrarCursosOfInstructor] error: " + res.Error.Error())
	}
	return inst
}

func (i *InstructorRepository) MostrarAllsInstructors() interface{} {
	return -1
}
