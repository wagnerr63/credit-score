package entities

type Debt struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Value    int    `json:"value"`
	Creditor string `json:"creditor"`
	Status   string `json:"status"`
}
