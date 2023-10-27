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
	if res := db.Orm.Model(new(models.Estudiante)).Delete(estudiante); res.Error != nil {
		panic("[Repositories.EstudianteRepository.EliminarEstudiante: Line 55] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (e *EstudianteRepository) EliminarEstudianteCurso(matricula *models.Matricula) interface{} {
	if res := db.Orm.Model(new(models.Matricula)).Delete(matricula); res.Error != nil {
		panic("[Repositories.EstudianteRepository.EliminarEstudianteCurso: Line 64] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (e *EstudianteRepository) MostrarEstudianteJuridicoNoJuridico(estudiante *models.Estudiante) interface{} {
	estudiantes := make([]models.Estudiante, 0)
	if res := db.Orm.Model(new(models.Estudiante)).Find(&estudiantes, estudiante); res.Error != nil {
		panic("[Repositories.EstudianteRepository.mostrarEstudianteJuridicoNoJuridico: Line 74] - " + res.Error.Error())
	}
	return estudiantes
}

func (e *EstudianteRepository) mostrarEstudianteDeCurso(estudiante *models.Estudiante) interface{} {
	return ""
}
