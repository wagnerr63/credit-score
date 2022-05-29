package people

import (
	"errors"

	"debts-service/entities"
)

type MockPeopleRepository struct {
	people []entities.Person `json:"people"`
	debts  []entities.Debt   `json:"debts"`
}

func NewMockPeopleRepository() IPeopleRepository {
	return &MockPeopleRepository{
		people: []entities.Person{
			{ID: 1, Name: "John", DocumentID: "123456789", Address: "Rio de Janeiro, Brazil", UserID: 1},
			{ID: 2, Name: "Mary", DocumentID: "987654321", Address: "Sao Paulo, Brazil", UserID: 2},
			{ID: 3, Name: "Bob", DocumentID: "123456789", Address: "Rio de Janeiro, Brazil", UserID: 3},
		},
		debts: []entities.Debt{
			{ID: 1, UserID: 1, Value: 100, Creditor: "Company A", Status: "Pending"},
			{ID: 2, UserID: 2, Value: 200, Creditor: "Company B", Status: "Pending"},
			{ID: 3, UserID: 3, Value: 300, Creditor: "Company C", Status: "Pending"},
		},
	}
}

func (r *MockPeopleRepository) FindByID(id int) (entities.Person, error) {
	for _, person := range r.people {
		if person.ID == id {
			return person, nil
		}
	}
	return entities.Person{}, errors.New("user not found")
}

func (r *MockPeopleRepository) FindPersonDebitsByPersonID(id int) (entities.PersonDebts, error) {
	var personDebits entities.PersonDebts

	for _, person := range r.people {
		if person.ID == id {
			personDebits.Person = person
			break
		}
	}

	for _, debt := range r.debts {
		if debt.UserID == id {
			personDebits.Debts = append(personDebits.Debts, debt)
		}
	}

	return personDebits, nil
}
