package infra

import (
	domain "github.com/fernandosvrosa/rinha-backend/api/src/domain/client"
	"github.com/fernandosvrosa/rinha-backend/api/src/primary/client"
	"github.com/fernandosvrosa/rinha-backend/api/src/secondary/client/adapter"
	"github.com/gocql/gocql"
)

type ClientFactory struct {
	session *gocql.Session
}

func NewClientFactory(session *gocql.Session) *ClientFactory {
	return &ClientFactory{
		session: session,
	}
}

func (cf *ClientFactory) CreateClientHandler() *client.ClientHandler {
	findClientAdapter := adapter.NewFindAccountAdapter(cf.session)
	updateAmountClientAdapter := adapter.NewUpdateAmountClientAdapter(cf.session)
	contaLockAdapter := adapter.NewContaLockAdapter(cf.session)
	contaUnLockAdapter := adapter.NewContaUnLockAdapter(cf.session)
	saveTransactionHistoryAdapter := adapter.NewSaveTransactionHistoryAdapter(cf.session)
	createTransactionUseCase := domain.NewCreateTransactionUsecase(
		findClientAdapter,
		updateAmountClientAdapter,
		contaLockAdapter,
		contaUnLockAdapter,
		saveTransactionHistoryAdapter,
	)
	return client.NewClientHandler(createTransactionUseCase)
}
