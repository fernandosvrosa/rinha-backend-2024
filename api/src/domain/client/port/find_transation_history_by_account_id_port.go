package port

import "github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"

type FindTransactionHistoryByAccountIdPort interface {
	Execute(accountID int) ([]entity.TransactionHistory, error)
}
