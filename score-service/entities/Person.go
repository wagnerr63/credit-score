package entities

type Person struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Age        int         `json:"age"`
	DocumentID string      `json:"document_id"`
	Belongings []Belonging `json:"belongings"`
}
