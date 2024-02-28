package adapter

import (
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"
	"github.com/gocql/gocql"
)

type SaveTransactionHistoryAdapter struct {
	session *gocql.Session
}

func NewSaveTransactionHistoryAdapter(session *gocql.Session) *SaveTransactionHistoryAdapter {
	return &SaveTransactionHistoryAdapter{session: session}
}

func (s SaveTransactionHistoryAdapter) Execute(history entity.TransactionHistory) (entity.TransactionHistory, error) {
	history.Id = gocql.TimeUUID()
	query := "INSERT INTO transaction_history (id, account_id, amount, type, description, created_at) VALUES (?, ?, ?, ?, ?, ?)"
	if err := s.session.Query(query, history.Id, history.AccountId, history.Amount, history.Type, history.Description, history.CreatedAt).Exec(); err != nil {
		return entity.TransactionHistory{}, err
	}
	return history, nil
}
