package repositories

import "ecs_govel/app/model"

type RepositoryTest struct {
	Repository
}

func NewRepositoryTemplate() *RepositoryTest {
	r := &RepositoryTest{}
	r.Repository = &RepositoryKernel{
		_type:  RepositoriesTypeTest,
		_rep:   r,
		_model: &model.TestModel{},
	}

	return r
}

// Implementacion de metodos del repositorio
func (*RepositoryTest) MetTest() {}
