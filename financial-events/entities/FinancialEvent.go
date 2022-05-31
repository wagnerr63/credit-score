package entities

type FinancialEvent struct {
	ID          int    `json:"id"`
	DocumentID  string `json:"document_id"`
	Description string `json:"description"`
	Value       int    `json:"value"`
	CreatedAt   string `json:"created_at"`
}
