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
	findClientAdapter := adapter.NewFindClientAdapter(cf.session)
	updateAmountClientAdapter := adapter.NewUpdateAmountClientAdapter(cf.session)
	createTransactionUseCase := domain.NewCreateTransactionUsecase(findClientAdapter, updateAmountClientAdapter)
	return client.NewClientHandler(createTransactionUseCase)
}
