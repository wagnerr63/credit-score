package repositories

import (
	"financial-events/repositories/financial_events"
)

type Container struct {
	FinancialEventRepository financial_events.IFinancialEventsRepository
}

func New() *Container {
	return &Container{
		FinancialEventRepository: financial_events.NewMockFinancialEventsRepository(),
	}
}
