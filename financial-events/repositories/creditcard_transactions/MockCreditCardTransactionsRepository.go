package creditcard_transactions

import "financial-events/entities"

type MockCrediCardTransactionsRepository struct {
	creditCardTransactions []entities.CreditCardTransaction
}

func NewMockCreditCardTransactionsRepository() ICreditCardTransactionsRepository {
	return &MockCrediCardTransactionsRepository{
		creditCardTransactions: []entities.CreditCardTransaction{
			{
				ID:          1,
				DocumentID:  "123456789",
				Description: "Spotify",
				Value:       100,
				CreatedAt:   "2020-01-01 00:00:00",
			},
			{
				ID:          2,
				DocumentID:  "123456789",
				Description: "Netflix",
				Value:       10000,
				CreatedAt:   "2020-01-01 00:00:00",
			},
		},
	}

}

func (m *MockCrediCardTransactionsRepository) FindLastByDocumentID(documentID string) (entities.CreditCardTransaction, error) {
	var creditCardTransaction entities.CreditCardTransaction
	for _, creditCardTransaction = range m.creditCardTransactions {
		if creditCardTransaction.DocumentID == documentID {
			return creditCardTransaction, nil
		}
	}

	return entities.CreditCardTransaction{}, nil
}
