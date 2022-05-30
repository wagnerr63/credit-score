package person

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

func (u *personUseCases) FindByDocumentID(documentID string, token string) (entities.Person, error) {

	postBody, _ := json.Marshal(ValidateRequestDTO{
		Action: "listPersonBelongings",
	})

	requestBody := bytes.NewBuffer(postBody)

	request, _ := http.NewRequest("POST", "http://localhost:8080/auth/validateRequest", requestBody)
	request.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": []string{"Bearer " + token},
	}

	fmt.Println("teste")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return entities.Person{}, errors.New("error trying validating request")
	}

	fmt.Println(resp)
	if resp.StatusCode != 200 {
		return entities.Person{}, errors.New("error validating request")
	}

	person, _ := u.repository.People.FindByDocumentID(documentID)

	if person.ID == 0 {
		return entities.Person{}, errors.New("person not found")
	}

	return person, nil
}
