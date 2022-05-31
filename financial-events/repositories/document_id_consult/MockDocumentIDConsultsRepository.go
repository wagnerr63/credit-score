package document_id_consult

import "financial-events/entities"

type MockDocumentIDConsultsRepository struct {
	documentIDConsults []entities.DocumentIDConsult
}

func NewMockDocumentIDConsultsRepository() IDocumentIDConsultsRepository {
	return &MockDocumentIDConsultsRepository{
		documentIDConsults: []entities.DocumentIDConsult{
			{
				ID:           1,
				DocumentID:   "123456789",
				CreditBureau: "Company A",
				ConsultDate:  "2020-01-01",
			},
			{
				ID:           2,
				DocumentID:   "123456789",
				CreditBureau: "Company B",
				ConsultDate:  "2020-01-01",
			},
			{
				ID:           3,
				DocumentID:   "223456789",
				CreditBureau: "Company C",
				ConsultDate:  "2020-01-01",
			},
		},
	}
}

func (m *MockDocumentIDConsultsRepository) FindLastConsultByDocumentID(documentID string) (entities.DocumentIDConsult, error) {
	for _, documentIDConsult := range m.documentIDConsults {
		if documentIDConsult.DocumentID == documentID {
			return documentIDConsult, nil
		}
	}
	return entities.DocumentIDConsult{}, nil
}
