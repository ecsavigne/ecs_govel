package repositories

type RepositoryTemplate struct {
	Repository
}

func NewRepositoryTemplate() *RepositoryTemplate {
	return &RepositoryTemplate{
		Repository: &RepositoryKernel{
			Type: RepositoriesTypeTemplate,
		},
	}
}

// Implementacion de metodos del repositorio
func (*RepositoryTemplate) MetTest() {}
