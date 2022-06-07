package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user := os.Getenv("user")
	dbname := os.Getenv("dbname")
	sslmode := "disable"
	password := os.Getenv("password")

	creds := fmt.Sprintf("user=%v dbname=%v sslmode=%v password=%v", user, dbname, sslmode, password)

	db, err := sql.Open("postgres", creds)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
