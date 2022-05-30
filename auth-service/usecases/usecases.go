package usecases

import (
	"auth-service/repositories"
	"auth-service/usecases/auth"
)

type Container struct {
	Auth auth.IAuthUseCases
}

type Options struct {
	Repositories *repositories.Container
}

func New(opt Options) *Container {
	return &Container{
		Auth: auth.New(opt.Repositories),
	}
}
