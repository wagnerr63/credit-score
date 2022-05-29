package people

import "score-service/entities"

type MockPeopleRepository struct {
	people []entities.Person `json:"people"`
}

func NewMockPeopleRepository() IPeopleRepository {
	return &MockPeopleRepository{
		people: []entities.Person{
			{ID: 1, Name: "John", Age: 30, DocumentID: "123456789", Belongings: []entities.Belonging{
				{ID: 1, Name: "Car", Value: 100000},
				{ID: 2, Name: "House", Value: 100000},
			}},
			{ID: 2, Name: "Jane", Age: 30, DocumentID: "223456789", Belongings: []entities.Belonging{
				{ID: 1, Name: "Car", Value: 100000},
				{ID: 2, Name: "House", Value: 100000},
			}},
		},
	}
}

func (m *MockPeopleRepository) FindByDocumentID(documentID string) (entities.Person, error) {
	for _, person := range m.people {
		if person.DocumentID == documentID {
			return person, nil
		}
	}
	return entities.Person{}, nil
}
