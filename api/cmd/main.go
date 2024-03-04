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

	cluster := gocql.NewCluster(dbHost)
	cluster.Keyspace = "rinha_db"
	cluster.Consistency = gocql.Quorum

	cluster.NumConns = 10
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())

	connManager := &infra.ConnectionManager{
		Cluster: cluster,
	}

	if err := connManager.Connect(); err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer connManager.Close()

	app := fiber.New()

	clientFactory := infra.NewClientFactory(connManager.Session)
	clientHandler := clientFactory.CreateClientHandler()

	transactionHistoryFactory := infra.NewTransactionHistoryFactory(connManager.Session)
	transactionHistoryHandler := transactionHistoryFactory.CreateTransactionHistoryHandler()

	app.Post("/clientes/:id/transacoes", clientHandler.CreateTransaction)
	app.Get("/clientes/:id/extrato", transactionHistoryHandler.FindTransactionHistory)

	log.Fatal(app.Listen(":3000"))
}
