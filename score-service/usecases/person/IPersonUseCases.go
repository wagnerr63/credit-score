package person

import "score-service/entities"

type IPersonUsecases interface {
	FindByDocumentID(documentID string) (entities.Person, error)
}
