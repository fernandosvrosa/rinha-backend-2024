package main

import (
	"github.com/fernandosvrosa/rinha-backend/api/infra"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "rinha_db"

	// Create a session
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Error creating session:", err)
	}
	defer session.Close()

	app := fiber.New()

	clientFactory := infra.NewClientFactory(session)
	clientHandler := clientFactory.CreateClientHandler()

	transactionHistoryFactory := infra.NewTransactionHistoryFactory(session)
	transactionHistoryHandler := transactionHistoryFactory.CreateTransactionHistoryHandler()

	app.Post("/clientes/:id/transacoes", clientHandler.CreateTransaction)
	app.Get("/clientes/:id/extrato", transactionHistoryHandler.FindTransactionHistory)

	log.Fatal(app.Listen(":3000"))
}
