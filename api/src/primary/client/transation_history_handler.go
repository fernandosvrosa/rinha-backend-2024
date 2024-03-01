package client

import (
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client"
	appError "github.com/fernandosvrosa/rinha-backend/api/src/domain/client/error"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"time"
)

type TransactionHistoryHandler struct {
	findTransactionHistoryUsecase *client.FindTransactionHistoryUsecase
	findAccountByIdUsecase        *client.FindAccountByIdUsecase
}

func NewTransactionHistoryHandler(
	findTransactionHistoryUsecase *client.FindTransactionHistoryUsecase,
	findAccountByIdUsecase *client.FindAccountByIdUsecase,
) *TransactionHistoryHandler {
	return &TransactionHistoryHandler{
		findTransactionHistoryUsecase: findTransactionHistoryUsecase,
		findAccountByIdUsecase:        findAccountByIdUsecase,
	}
}

type (
	Account struct {
		Amount int64     `json:"total"`
		Date   time.Time `json:"data_extrato"`
		Limit  int64     `json:"limite"`
	}

	Transaction struct {
		Amount      int64     `json:"valor"`
		Type        string    `json:"tipo"`
		Description string    `json:"descricao"`
		CreatedAt   time.Time `json:"realizada_em"`
	}

	TransactionHistoryResponse struct {
		Account      Account       `json:"saldo"`
		Transactions []Transaction `json:"ultimas_transacoes"`
	}
)

func (t *TransactionHistoryHandler) FindTransactionHistory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	account, err := t.findAccountByIdUsecase.Execute(id)

	if err != nil {
		switch err.(type) {
		case appError.NotFound:
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	transactionsHistory, err := t.findTransactionHistoryUsecase.Execute(id)

	if err != nil {
		switch err.(type) {
		case appError.NotFound:
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	accountResponse := Account{
		Amount: account.Amount,
		Date:   time.Now(),
		Limit:  account.Limit,
	}

	transactionsResponse := make([]Transaction, len(transactionsHistory))

	for i, transaction := range transactionsHistory {
		transactionsResponse[i] = Transaction{
			Amount:      transaction.Amount,
			Type:        transaction.Type,
			Description: transaction.Description,
			CreatedAt:   transaction.CreatedAt,
		}
	}

	err = c.Status(http.StatusOK).JSON(TransactionHistoryResponse{
		Account:      accountResponse,
		Transactions: transactionsResponse,
	})
	if err != nil {
		return err
	}
	return nil
}
