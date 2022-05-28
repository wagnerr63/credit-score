package users

import "auth-service/entities"

type MockUsersRepository struct {
	users []entities.User
}

func NewMockUsersRepository() IUsersRepository {
	return &MockUsersRepository{users: []entities.User{
		{ID: 1, Email: "john@mail.com", Password: "123456", Role: "client"},
		{ID: 2, Email: "marie@mail.com", Password: "123456", Role: "client"},
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
