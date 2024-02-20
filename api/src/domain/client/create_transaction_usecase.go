package client

import (
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"
	appError "github.com/fernandosvrosa/rinha-backend/api/src/domain/client/error"
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/port"
)

type CreateTransactionUsecase struct {
	findClientPort         port.FindClientPort
	updateAmountClientPort port.UpdateAmountClientPort
}

func NewCreateTransactionUsecase(findClientPort port.FindClientPort, updateAmountClientPort port.UpdateAmountClientPort) *CreateTransactionUsecase {
	return &CreateTransactionUsecase{findClientPort: findClientPort, updateAmountClientPort: updateAmountClientPort}
}

func (c *CreateTransactionUsecase) Execute(transaction entity.Transaction) (entity.Balance, error) {

	client, err := c.findClientPort.Execute(transaction.ClientID)

	if err != nil {
		return entity.Balance{}, err
	}

	client, err = balanceAction(client, transaction)

	if err != nil {
		return entity.Balance{}, err
	}

	client, err = c.updateAmountClientPort.Execute(client)

	if err != nil {
		return entity.Balance{}, err
	}

	return entity.Balance{
		Amount: client.Amount,
		Limit:  client.Limit,
	}, nil
}

func balanceAction(client entity.Client, transaction entity.Transaction) (entity.Client, error) {
	if transaction.TransactionType == "c" {
		client.Amount += transaction.Value
	} else {
		client.Amount -= transaction.Value
	}

	if client.Limit < client.Amount {
		return entity.Client{}, appError.InsufficientFund{Message: "Insufficient fund."}
	}

	return client, nil
}
