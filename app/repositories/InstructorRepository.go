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
	if res := db.Orm.Model(new(models.Persona)).Updates(instructor); res.Error != nil {
		panic("[Repositories.InstructorRepository.ModificarInstructor: Line 24] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (i *InstructorRepository) EliminarInstructor(instructor *models.Instructore) interface{} {
	if res := db.Orm.Model(new(models.Instructore)).Delete(instructor); res.Error != nil {
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

func (i *InstructorRepository) AgregarInstructorToCurso(instructor_curso *models.InstructorCurso) interface{} {
	if err := db.Orm.Create(instructor_curso).Error; err != nil {
		panic("[Repositories.InstructorRepository.AgregarInstructorToCurso: Line 51] - " + err.Error())
	}
	return true
}

func (i *InstructorRepository) EliminarCursoOffInstructor(instructor_curso *models.InstructorCurso) interface{} {
	if res := db.Orm.Model(new(models.InstructorCurso)).Delete(instructor_curso); res.Error != nil {
		panic("[Repositories.InstructorRepository.EliminarInstructorDeCurso: Line 58] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (i *InstructorRepository) ModificarInstructorCurso(instructor_curso *models.InstructorCurso) interface{} {
	if res := db.Orm.Model(new(models.InstructorCurso)).Updates(instructor_curso); res.Error != nil {
		panic("[Repositories.InstructorRepository.ModificarInstructorCurso: Line 67] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (i *InstructorRepository) MostrarAllsInstructors() interface{} {
	return -1
}
