package creditcard_transactions

import (
	"financial-events/entities"
)

type ICreditCardTransactionsRepository interface {
	FindLastByDocumentID(documentID string) (entities.CreditCardTransaction, error)
}
