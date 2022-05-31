package financial_events

import "financial-events/entities"

type IFinancialEventsRepository interface {
	ListByDocumentID(documentID string) ([]entities.FinancialEvent, error)
}
