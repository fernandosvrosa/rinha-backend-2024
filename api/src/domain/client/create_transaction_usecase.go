package client

import (
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"
	appError "github.com/fernandosvrosa/rinha-backend/api/src/domain/client/error"
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/port"
	"time"
)

type CreateTransactionUsecase struct {
	findAccountPort            port.FindAccountPort
	updateAmountClientPort     port.UpdateAmountClientPort
	contaLockPort              port.ContaLockPort
	contaUnLockPort            port.ContaUnLockPort
	saveTransactionHistoryPort port.SaveTransactionHistoryPort
}

func NewCreateTransactionUsecase(
	findAccountPort port.FindAccountPort,
	updateAmountClientPort port.UpdateAmountClientPort,
	contaLockPort port.ContaLockPort,
	contaUnLockPort port.ContaUnLockPort,
	saveTransactionHistoryPort port.SaveTransactionHistoryPort,
) *CreateTransactionUsecase {
	return &CreateTransactionUsecase{
		findAccountPort:            findAccountPort,
		updateAmountClientPort:     updateAmountClientPort,
		contaLockPort:              contaLockPort,
		contaUnLockPort:            contaUnLockPort,
		saveTransactionHistoryPort: saveTransactionHistoryPort,
	}
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

	account, err := c.findAccountPort.Execute(transaction.ClientID)

	if err != nil {
		if err.Error() == "not found" {
			return entity.Balance{}, appError.NotFound{Message: "Account not found."}

		}
		return entity.Balance{}, err
	}

	account, err = balanceAction(account, transaction)

	if err != nil {
		return entity.Balance{}, err
	}

	account, err = c.updateAmountClientPort.Execute(account)

	if err != nil {
		return entity.Balance{}, err
	}

	transactionHistory := entity.TransactionHistory{
		AccountId:   account.ID,
		CreatedAt:   time.Now(),
		Amount:      transaction.Value,
		Description: transaction.Description,
		Type:        transaction.TransactionType,
	}

	_, err = c.saveTransactionHistoryPort.Execute(transactionHistory)

	if err != nil {
		return entity.Balance{}, err
	}

	return entity.Balance{
		Amount: account.Amount,
		Limit:  account.Limit,
	}, nil
}

func balanceAction(client entity.Account, transaction entity.Transaction) (entity.Account, error) {
	if transaction.TransactionType == "d" {
		client.Amount += transaction.Value
	} else {
		client.Amount -= transaction.Value
	}

	if client.Amount < 0 {
		if client.Limit < -client.Amount {
			return entity.Account{}, appError.InsufficientFund{Message: "Insufficient fund."}
		}
	}

	return client, nil
}
