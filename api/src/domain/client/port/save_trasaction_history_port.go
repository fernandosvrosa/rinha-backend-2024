package port

import "github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"

type SaveTransactionHistoryPort interface {
	Execute(transaction entity.TransactionHistory) (entity.TransactionHistory, error)
}
