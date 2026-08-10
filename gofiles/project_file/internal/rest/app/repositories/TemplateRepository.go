package repositories

import "ecs_govel/internal/rest/app/model"

type TestRepository struct {
	*KernelRepository
}

func NewRepositoryTemplate() *TestRepository {
	r := new(TestRepository)

	kernel := &KernelRepository{
		_type:  RepositoriesTypeTest,
		_rep:   r,
		_model: &model.TestModel{},
	}
	r.KernelRepository = kernel

	return r
}

// Implementacion de metodos del repositorio
func (*TestRepository) MetTest() {}
