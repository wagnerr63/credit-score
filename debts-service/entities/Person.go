package entities

type Person struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	DocumentID string `json:"document_id"`
	Address    string `json:"address"`
	UserID     int    `json:"user_id"`
}
