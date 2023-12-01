package repositories

import (
	"ecs_govel/config/db"
)

type MatriculaRepository struct {
}

// Metodos del Repositorio
// se le pasa una matricula con id de curso, muestra matriculas de un curso
func (m *MatriculaRepository) MostrarMatricula(idCurso int) interface{} {
	matriculas := []map[string]interface{}{}
	if estuds := db.Orm.
		Table("cursos").
		Select("cursos.id as idCurso, personas.nombre as nombreEst, personas.apellidos as apellidoEst,"+
			"estudiantes.ci as est_ci, personas.mail as mailEst, personas.dir as dirEst,"+
			"estudiantes.si_juridico as jurid").
		Joins("inner join matriculas on matriculas.curso_id = cursos.id").
		Joins("inner join estudiantes on matriculas.estudiante_ci = estudiantes.ci").
		Joins("inner join personas on estudiantes.ci = personas.ci").
		Where("cursos.id = ?", idCurso).
		Order("personas.nombre asc"); estuds.Error != nil {
		panic("Ocurried one error in line 23 [MatriculaRepository.MostrarMatricula] error: " + estuds.Error.Error())
	} else {
		if insts := db.Orm.
			Table("cursos").
			Select("cursos.id as idCursoIns, personas.nombre as nombreInst, personas.apellidos as apellidoInst,"+
				"instructores.ci as inst_ci, personas.mail as mailInst").
			Joins("inner join instructor_cursos on cursos.id = instructor_cursos.curso_id").
			Joins("inner join instructores on instructores.ci = instructor_cursos.instructor_ci").
			Joins("inner join personas on personas.ci = instructores.ci").
			Where("cursos.id = ?", idCurso).
			Order("personas.nombre asc"); insts.Error != nil {
			panic("Ocurried one error in line 34 [MatriculaRepository.MostrarMatricula] error: " + insts.Error.Error())
		} else {
			if res := db.Orm.Table("cursos").Joins("inner join (?) est on est.idCurso = cursos.id", estuds).
				Joins("inner join (?) inst on inst.idCursoIns = cursos.id", insts).
				Find(&matriculas); res.Error != nil {
				panic("Ocurried one error in line 38 [MatriculaRepository.MostrarMatricula] error: " + res.Error.Error())
			}
		}
	}
	return matriculas
}
