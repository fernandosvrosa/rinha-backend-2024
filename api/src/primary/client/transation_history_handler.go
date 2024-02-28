package client

import (
	"fmt"
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client"
	appError "github.com/fernandosvrosa/rinha-backend/api/src/domain/client/error"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type TransactionHistoryHandler struct {
	findTransactionHistoryUsecase client.FindTransactionHistoryUsecase
	findAccountByIdUsecase        client.FindAccountByIdUsecase
}

func NewTransactionHistoryHandler(
	findTransactionHistoryUsecase client.FindTransactionHistoryUsecase,
	findAccountByIdUsecase client.FindAccountByIdUsecase,
) *TransactionHistoryHandler {
	return &TransactionHistoryHandler{
		findTransactionHistoryUsecase: findTransactionHistoryUsecase,
		findAccountByIdUsecase:        findAccountByIdUsecase,
	}
}

func (t *TransactionHistoryHandler) SaveTransactionHistory(c *fiber.Ctx) error {
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

	fmt.Println("account: ", account)

	transactionsHistory, err := t.findTransactionHistoryUsecase.Execute(id)

	if err != nil {
		switch err.(type) {
		case appError.NotFound:
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	fmt.Println("transactionsHistory: ", transactionsHistory)

	return nil
}
