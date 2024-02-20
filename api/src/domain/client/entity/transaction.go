package entity

type Transaction struct {
	ClientID        int
	Value           int64
	TransactionType string
	Description     string
}
