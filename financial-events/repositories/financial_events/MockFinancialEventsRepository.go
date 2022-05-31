package financial_events

import "financial-events/entities"

type MockFinancialEventsRepository struct {
	financialEvents []entities.FinancialEvent
}

func NewMockFinancialEventsRepository() IFinancialEventsRepository {
	return &MockFinancialEventsRepository{
		financialEvents: []entities.FinancialEvent{
			{
				ID:          1,
				DocumentID:  "123456789",
				Description: "transfer",
				Value:       100,
				CreatedAt:   "2018-01-01 00:00:00",
			},
			{
				ID:          2,
				DocumentID:  "123456789",
				Description: "deposit",
				Value:       200,
				CreatedAt:   "2018-01-01 00:00:00",
			},
			{
				ID:          3,
				DocumentID:  "223456789",
				Description: "transfer",
				Value:       100,
				CreatedAt:   "2018-01-01 00:00:00",
			},
		},
	}

}

func (repository *MockFinancialEventsRepository) ListByDocumentID(documentID string) ([]entities.FinancialEvent, error) {
	var financialEvents []entities.FinancialEvent
	for _, financialEvent := range repository.financialEvents {
		if financialEvent.DocumentID == documentID {
			financialEvents = append(financialEvents, financialEvent)
		}
	}
	return financialEvents, nil
}
