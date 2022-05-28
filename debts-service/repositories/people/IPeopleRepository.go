package people

import (
	"debts-service/entities"
)

type IPeopleRepository interface {
	FindByID(id int) (entities.Person, error)
	FindPersonDebitsByPersonID(id int) (entities.PersonDebts, error)
}
