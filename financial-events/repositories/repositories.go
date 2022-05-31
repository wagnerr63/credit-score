package repositories

import (
	"financial-events/repositories/creditcard_transactions"
	"financial-events/repositories/document_id_consult"
	"financial-events/repositories/financial_events"
)

type Container struct {
	FinancialEventRepository        financial_events.IFinancialEventsRepository
	DocumentIDConsultRepository     document_id_consult.IDocumentIDConsultsRepository
	CreditCardTransactionRepository creditcard_transactions.ICreditCardTransactionsRepository
}

func New() *Container {
	return &Container{
		FinancialEventRepository:        financial_events.NewMockFinancialEventsRepository(),
		DocumentIDConsultRepository:     document_id_consult.NewMockDocumentIDConsultsRepository(),
		CreditCardTransactionRepository: creditcard_transactions.NewMockCreditCardTransactionsRepository(),
	}
}
