package usecases

import (
	"financial-events/repositories"
)

type Container struct{}

type Options struct {
	Repositories *repositories.Container
}

func New(opt Options) *Container {
	return &Container{}
}
