package people

import (
	"score-service/entities"
)

type IPeopleRepository interface {
	FindByDocumentID(documentID string) (entities.Person, error)
}
