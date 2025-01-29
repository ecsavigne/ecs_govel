package repositories

import "ecs_govel/app/model"


type RepositoryType = string

const (
	RepositoriesTypeTemplate RepositoryType = "repository_type_template"
	// Asi crear el resto
)

type Repository interface {
	// Funciones comunes para todos los repositorios que embeben Repository y se implementan en RepositoryKernelssss
	GetType() string
	GetRepository() Repository
}

type RepositoryKernel struct {
	_type string
	_rep   Repository
	_model model.ModelKernel
}

func (r *RepositoryKernel) GetType() string {
	return r.Type
}

func (r *RepositoryKernel) GetRepository() Repository {
	return r._rep
}

func (r *RepositoryKernel) GetModel() model.ModelKernel {
	return r._model
}
