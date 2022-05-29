package usecases

import (
	"score-service/repositories"
	"score-service/usecases/person"
)

type Container struct {
	Person person.IPersonUsecases
}

type Options struct {
	Repositories *repositories.Container
}

func New(options Options) *Container {
	return &Container{
		Person: person.New(options.Repositories),
	}
}
