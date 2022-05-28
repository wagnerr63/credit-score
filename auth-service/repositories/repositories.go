package repositories

import "auth-service/repositories/users"

type Container struct {
	User users.IUsersRepository
}

func New() *Container {
	return &Container{
		User: users.NewMockUsersRepository(),
	}
}
