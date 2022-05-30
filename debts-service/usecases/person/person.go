package person

import (
	"bytes"
	"debts-service/entities"
	"debts-service/repositories"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type IPersonUseCases interface {
	GetPersonDebtsByPersonID(id int, token string) (entities.PersonDebts, error)
}

type personUseCases struct {
	repositories *repositories.Container
}

type ValidateRequestDTO struct {
	Action string `json:"action"`
	UserID int    `json:"user_id"`
}

func New(repo *repositories.Container) IPersonUseCases {
	return &personUseCases{
		repositories: repo,
	}
}

func (u *personUseCases) GetPersonDebtsByPersonID(id int, token string) (entities.PersonDebts, error) {
	personByID, err := u.repositories.People.FindPersonDebitsByPersonID(id)
	if err != nil {
		return entities.PersonDebts{}, err
	}

	postBody, _ := json.Marshal(ValidateRequestDTO{
		Action: "listPersonDebts",
		UserID: personByID.UserID,
	})

	requestBody := bytes.NewBuffer(postBody)

	request, err := http.NewRequest("POST", "http://localhost:8080/auth/validateRequest", requestBody)
	request.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": []string{"Bearer " + token},
	}

	resp, err := http.DefaultClient.Do(request)
	fmt.Println(resp)

	if resp.StatusCode != 200 {
		return entities.PersonDebts{}, errors.New("error validating request")
	}

	return u.repositories.People.FindPersonDebitsByPersonID(id)
}
