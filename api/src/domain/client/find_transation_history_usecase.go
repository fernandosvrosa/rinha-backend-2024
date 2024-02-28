package client

import "github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"

type FindTransactionHistoryUsecase struct {
}

func NewFindTransactionHistoryUsecase() *FindTransactionHistoryUsecase {
	return &FindTransactionHistoryUsecase{}
}

func (f *FindTransactionHistoryUsecase) Execute(accountId int) ([]entity.TransactionHistory, error) {

	return []entity.TransactionHistory{}, nil
}
