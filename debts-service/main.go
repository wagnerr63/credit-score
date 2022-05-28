package main

import (
	"debts-service/repositories"
	"debts-service/usecases"
	"fmt"
)

func main() {
	repositories := repositories.New()

	usecases := usecases.New(usecases.Options{
		Repositories: repositories,
	})

	fmt.Println(usecases.Person.GetPersonDebtsByPersonID(1))
}
