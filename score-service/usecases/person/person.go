package person

import (
	"score-service/entities"
	"score-service/repositories"
)

type personUseCases struct {
	repository *repositories.Container
}

func New(repo *repositories.Container) IPersonUsecases {
	return &personUseCases{
		repository: repo,
	}
}

func (u *personUseCases) FindByDocumentID(documentID string) (entities.Person, error) {
	return u.repository.People.FindByDocumentID(documentID)
}
