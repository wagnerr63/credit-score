package person

import (
	"debts-service/entities"
	"debts-service/repositories"
)

type IPersonUseCases interface {
	GetPersonDebtsByPersonID(id int) (entities.PersonDebts, error)
}

type personUseCases struct {
	repositories *repositories.Container
}

func New(repo *repositories.Container) IPersonUseCases {
	return &personUseCases{
		repositories: repo,
	}
}

func (u *personUseCases) GetPersonDebtsByPersonID(id int) (entities.PersonDebts, error) {
	// Validate token
	return u.repositories.People.FindPersonDebitsByPersonID(id)
}
