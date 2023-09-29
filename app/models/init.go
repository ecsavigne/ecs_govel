package models

// Arreglo que contendra todas las modelos
var Models []interface{}

func init() {
	Models = make([]interface{}, 0)
	// Agregado de Modelos
	Models = append(Models, &Persona{})
	Models = append(Models, &Matricula{})
	Models = append(Models, &Instructore{})
	Models = append(Models, &InstructorCurso{})
	Models = append(Models, &Estudiante{})
	Models = append(Models, &Estructura{})
	Models = append(Models, &EstructuraContenido{})
	Models = append(Models, &Curso{})
	Models = append(Models, &CursoContenido{})
	Models = append(Models, &Contenido{})
	Models = append(Models, &Usuario{})
}
