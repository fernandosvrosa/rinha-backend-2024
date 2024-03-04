package main

import (
	"fmt"
	"github.com/fernandosvrosa/rinha-backend/api/infra"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	dbHost := os.Getenv("DB_HOST")
	name := os.Getenv("NAME")

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>> NAME:", name)

	cluster := gocql.NewCluster(dbHost)
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
