package users

import (
	"auth-service/entities"
)

type IUsersRepository interface {
	FindByEmail(email string) (entities.User, error)
}
