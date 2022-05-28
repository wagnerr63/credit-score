package usecases

import (
	"debts-service/repositories"
	"debts-service/usecases/person"
)

type Container struct {
	Person person.IPersonUseCases
}

type Options struct {
	Repositories *repositories.Container
}

func New(opt Options) *Container {
	return &Container{
		Person: person.New(opt.Repositories),
	}
}
