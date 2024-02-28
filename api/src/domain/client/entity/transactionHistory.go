package entity

import (
	"github.com/gocql/gocql"
	"time"
)

type TransactionHistory struct {
	AccountId   int
	Id          gocql.UUID
	CreatedAt   time.Time
	Amount      int64
	Description string
	Type        string
}
