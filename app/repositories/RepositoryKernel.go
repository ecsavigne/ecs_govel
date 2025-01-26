package repositories

type Repository interface {
	// Funciones comunes para todos los repositorios que embeben Repository y se implementan en RepositoryKernelssss
	GetType() string
}

type RepositoryKernel struct {
	Type string
}

func (r *RepositoryKernel) GetType() string {
	return r.Type
}

const (
	RepositoriesTypeTemplate = "repository_type_template"
)
