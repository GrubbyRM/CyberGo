package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/joho/godotenv"         // For loading .env files
)

var db *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Could not connect to the databases: %v", err)
	}

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not pinging the database: %v", err)
	}

	fmt.Println("Connected successfully!")
}

func GetDB() *sql.DB {
	return db
}
