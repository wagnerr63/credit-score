package entities

type DocumentIDConsult struct {
	Id           int    `json:"id"`
	DocumentID   string `json:"document_id"`
	CreditBureau string `json:"credit_bureau"`
	CreatedAt    string `json:"created_at"`
}
