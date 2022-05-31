package document_id_consult

import "financial-events/entities"

type IDocumentIDConsultsRepository interface {
	FindLastConsultByDocumentID(documentID string) (entities.DocumentIDConsult, error)
}
