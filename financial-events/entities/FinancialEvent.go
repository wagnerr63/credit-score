package entities

type FinancialEvent struct {
	Id          int    `json:"id"`
	DocumentID  string `json:"document_id"`
	Description string `json:"description"`
	Value       string `json:"value"`
	CreatedAt   string `json:"created_at"`
}
