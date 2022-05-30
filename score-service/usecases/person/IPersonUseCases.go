package person

import "score-service/entities"

type IPersonUsecases interface {
	FindByDocumentID(documentID string, token string) (entities.Person, error)
}

type ValidateRequestDTO struct {
	Action string `json:"action"`
	UserID int    `json:"user_id"`
}
