package infra

import (
	client2 "github.com/fernandosvrosa/rinha-backend/api/src/domain/client"
	"github.com/fernandosvrosa/rinha-backend/api/src/primary/client"
	"github.com/fernandosvrosa/rinha-backend/api/src/secondary/client/adapter"
	"github.com/gocql/gocql"
)

type TransactionHistoryFactory struct {
	session *gocql.Session
}

func NewTransactionHistoryFactory(session *gocql.Session) *TransactionHistoryFactory {
	return &TransactionHistoryFactory{
		session: session,
	}
}

func (tf *TransactionHistoryFactory) CreateTransactionHistoryHandler() *client.TransactionHistoryHandler {
	findTransactionHistoryByAccountId := adapter.NewFindTransactionHistoryByAccountIdAdapter(tf.session)
	findTransactionHistoryUsecase := client2.NewFindTransactionHistoryUsecase(findTransactionHistoryByAccountId)

	findAccount := adapter.NewFindAccountAdapter(tf.session)
	findAccountByIdUsecase := client2.NewFindAccountByIdUsecase(findAccount)
	return client.NewTransactionHistoryHandler(findTransactionHistoryUsecase, findAccountByIdUsecase)
}
