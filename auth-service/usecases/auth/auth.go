package auth

import (
	"auth-service/repositories"
	"errors"
)

type authUseCases struct {
	repositories *repositories.Container
}

func New(repo *repositories.Container) IAuthUseCases {
	return &authUseCases{
		repositories: repo,
	}

}

func (a *authUseCases) ValidateRequest(request ValidateRequestDTO) error {
	if request.Action == "listPersonDebts" {
		return ValidateListPersonDebtsRequest(ValidateListPersonDebtsRequestDTO{
			UserID: request.UserID,
			Token:  request.Token,
			Role:   request.Role,
		})
	}

	return errors.New("invalid action")
}
