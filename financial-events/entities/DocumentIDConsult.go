package entities

type DocumentIDConsult struct {
	ID           int    `json:"id"`
	DocumentID   string `json:"document_id"`
	CreditBureau string `json:"credit_bureau"`
	ConsultDate  string `json:"consult_date"`
}
