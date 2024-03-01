package client

import (
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/port"
)

type FindTransactionHistoryUsecase struct {
	findTransactionHistoryByAccountId port.FindTransactionHistoryByAccountIdPort
}

func NewFindTransactionHistoryUsecase(findTransactionHistoryByAccountId port.FindTransactionHistoryByAccountIdPort) *FindTransactionHistoryUsecase {
	return &FindTransactionHistoryUsecase{findTransactionHistoryByAccountId: findTransactionHistoryByAccountId}
}

func (f *FindTransactionHistoryUsecase) Execute(accountId int) ([]entity.TransactionHistory, error) {
	transaction, err := f.findTransactionHistoryByAccountId.Execute(accountId)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
