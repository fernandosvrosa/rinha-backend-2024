package adapter

import (
	"fmt"
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"
	"github.com/gocql/gocql"
	"time"
)

type FindTransactionHistoryByAccountIdAdapter struct {
	session *gocql.Session
}

func NewFindTransactionHistoryByAccountIdAdapter(session *gocql.Session) *FindTransactionHistoryByAccountIdAdapter {
	return &FindTransactionHistoryByAccountIdAdapter{session: session}
}

func (f FindTransactionHistoryByAccountIdAdapter) Execute(accountID int) ([]entity.TransactionHistory, error) {
	query := fmt.Sprintf("SELECT id, account_id, amount, type, description, created_at FROM transaction_history WHERE account_id = %d LIMIT 10", accountID)

	iter := f.session.Query(query).Iter()

	var transactions []entity.TransactionHistory

	var transactionId gocql.UUID
	var accountId int
	var amount int64
	var transactionType string
	var description string
	var createdAt time.Time

	for iter.Scan(&transactionId, &accountId, &amount, &transactionType, &description, &createdAt) {
		transactions = append(transactions, entity.TransactionHistory{
			Id:          transactionId,
			AccountId:   accountID,
			Amount:      amount,
			Type:        transactionType,
			Description: description,
			CreatedAt:   createdAt,
		})
	}

	return transactions, nil
}
