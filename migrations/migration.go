// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"os"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/golang-migrate/migrate/v4"
// 	"github.com/golang-migrate/migrate/v4/database/mysql"
// 	_ "github.com/golang-migrate/migrate/v4/source/file"
// )

// func main() {
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbUser := os.Getenv("DB_USER")
// 	dbPassword := os.Getenv("DB_PASSWORD")
// 	dbName := os.Getenv("DB_NAME")

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
// 		dbUser, dbPassword, dbHost, dbPort, dbName)

// 	db, _ := sql.Open("mysql", dsn)
// 	driver, _ := mysql.WithInstance(db, &mysql.Config{})
// 	m, _ := migrate.NewWithDatabaseInstance(
// 		"file:///migrations",
// 		"mysql",
// 		driver,
// 	)

// 	m.Steps(2)
// }
