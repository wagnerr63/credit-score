package main

import (
	"fmt"
	"score-service/repositories"
	"score-service/usecases"
)

func main() {
	repositories := repositories.New()

	usecases := usecases.New(usecases.Options{
		Repositories: repositories,
	})

	person, err := usecases.Person.FindByDocumentID("123456789")
	if err != nil {
		panic(err)
	}

	fmt.Println(person)
}
