package repositories

import "ecs_govel/app/model"

type TestRepository struct {
	Repository
}

func NewRepositoryTemplate() *TestRepository {
	r := &TestRepository{}
	r.Repository = &KernelRepository{
		_type:  RepositoriesTypeTest,
		_rep:   r,
		_model: &model.TestModel{},
	}

	return r
}

// Implementacion de metodos del repositorio
func (*TestRepository) MetTest() {}
