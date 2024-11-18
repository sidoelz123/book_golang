package connection

import (
	"database/sql"
	"fmt"
	"os"
	"quiz-go/databases/migration"
	"quiz-go/databases/seeder"

	"github.com/gin-gonic/gin"
)

var (
	DBConnections *sql.DB
	err           error
)

func Initiator() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	env := os.Getenv("GO_ENV")

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	DBConnections, err = sql.Open("postgres", dbURL)

	// check connection
	err = DBConnections.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
	migration.DBMigrate(DBConnections)
	seeder.SeedUsers(DBConnections)
}
