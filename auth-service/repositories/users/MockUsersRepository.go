package users

import "auth-service/entities"

type MockUsersRepository struct {
	users []entities.User
}

func NewMockUsersRepository() IUsersRepository {
	return &MockUsersRepository{users: []entities.User{
		{ID: 1, Email: "john@mail.com", Password: "123456", Role: "client"},
		{ID: 2, Email: "marie@mail.com", Password: "123456", Role: "client"},
		{ID: 3, Email: "mike_credit_institution_A@mail.com", Password: "123456", Role: "credit_institution"},
		{ID: 4, Email: "credit_analyst@mail.com", Password: "123456", Role: "credit_analyst"},
	}}
}

func (r *MockUsersRepository) FindByEmail(email string) (entities.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return entities.User{}, nil
}

func (r *MockUsersRepository) FindByID(id int) (entities.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return entities.User{}, nil
}
