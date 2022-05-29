package repositories

import "score-service/repositories/people"

type Container struct {
	People people.IPeopleRepository
}

func New() *Container {
	return &Container{
		People: people.NewMockPeopleRepository(),
	}
}
