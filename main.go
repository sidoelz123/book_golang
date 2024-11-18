package main

import (
	"database/sql"
	"log"
	"quiz-go/databases/connection"
	"quiz-go/routes"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println(".env file loaded successfully")
	}

	connection.Initiator()
	defer connection.DBConnections.Close()
	routes.InitRouter()
}
