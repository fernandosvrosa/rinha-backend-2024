package client

import (
	"fmt"
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client"
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"
	appError "github.com/fernandosvrosa/rinha-backend/api/src/domain/client/error"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type ClientHandler struct {
	createTransactionUseCase *client.CreateTransactionUsecase
}

func NewClientHandler(createTransactionUseCase *client.CreateTransactionUsecase) *ClientHandler {
	return &ClientHandler{createTransactionUseCase: createTransactionUseCase}
}

type (
	Resquest struct {
		Value           int64  `json:"valor"`
		TransactionType string `json:"tipo"`
		Description     string `json:"descricao"`
	}

	ClientResponse struct {
		Limit  int64 `json:"limite"`
		Amount int64 `json:"saldo"`
	}
)

func (ch *ClientHandler) CreateTransaction(c *fiber.Ctx) error {
	fmt.Println("CreateTransaction")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	request := &Resquest{}
	if err := c.BodyParser(request); err != nil {
		c.Status(400).JSON(fiber.Map{"error": err.Error()})
		return nil
	}

	balance, err := ch.createTransactionUseCase.Execute(
		entity.Transaction{
			ClientID:        id,
			Value:           request.Value,
			TransactionType: request.TransactionType,
			Description:     request.Description,
		})

	if err != nil {
		switch err.(type) {
		case appError.InsufficientFund:
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
		case appError.NotFound:
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	err = c.Status(http.StatusOK).JSON(ClientResponse{
		Limit:  balance.Limit,
		Amount: balance.Amount,
	})
	if err != nil {
		return err
	}
	return nil
}
