package repositories

import "ecs_govel/internal/rest/app/model"

type RepositoryType = string

const (
	RepositoriesTypeTemplate RepositoryType = "repository_type_template"
	RepositoriesTypeTest     RepositoryType = "repository_type_test"
	// Asi crear el resto
)

type Repository interface {
	// Funciones comunes para todos los repositorios que embeben Repository y se implementan en KernelRepositoryssss
	GetType() string
	GetRepository() Repository
}

type KernelRepository struct {
	_type  string
	_rep   Repository
	_model model.ModelKernel
}

func (r *KernelRepository) GetType() string {
	return r._type
}

func (r *KernelRepository) GetRepository() Repository {
	return r._rep
}

func (r *KernelRepository) GetModel() model.ModelKernel {
	return r._model
}
