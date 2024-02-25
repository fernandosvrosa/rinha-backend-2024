package client

import (
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"
	appError "github.com/fernandosvrosa/rinha-backend/api/src/domain/client/error"
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/port"
)

type CreateTransactionUsecase struct {
	findClientPort         port.FindClientPort
	updateAmountClientPort port.UpdateAmountClientPort
	contaLockPort          port.ContaLockPort
	contaUnLockPort        port.ContaUnLockPort
}

func NewCreateTransactionUsecase(
	findClientPort port.FindClientPort,
	updateAmountClientPort port.UpdateAmountClientPort,
	contaLockPort port.ContaLockPort,
	contaUnLockPort port.ContaUnLockPort,
) *CreateTransactionUsecase {
	return &CreateTransactionUsecase{
		findClientPort:         findClientPort,
		updateAmountClientPort: updateAmountClientPort,
		contaLockPort:          contaLockPort,
		contaUnLockPort:        contaUnLockPort}
}

func (c *CreateTransactionUsecase) Execute(transaction entity.Transaction) (entity.Balance, error) {
	applied, err := c.contaLockPort.Execute(transaction.ClientID)
	if err != nil {
		return entity.Balance{}, err
	}

	if !applied {
		c.Execute(transaction)
	}
	defer c.contaUnLockPort.Execute(transaction.ClientID)

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

	if client.Amount < 0 {
		if client.Limit < -client.Amount {
			return entity.Client{}, appError.InsufficientFund{Message: "Insufficient fund."}
		}
	}

	return client, nil
}
