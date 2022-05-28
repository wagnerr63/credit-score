package entities

type PersonDebts struct {
	Person
	Debts []Debt `json:"debts"`
}
