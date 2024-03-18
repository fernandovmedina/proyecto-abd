package database

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func OpenDB() error {
	var err error

	if err = godotenv.Load(); err != nil {
		return err
	}

	var (
		databaseName string = os.Getenv("DATABASE_NAME")
		databaseUser string = os.Getenv("DATABASE_USER")
		databasePass string = os.Getenv("DATABASE_PASS")
		databaseHost string = os.Getenv("DATABASE_HOST")
		databasePort string = os.Getenv("DATABASE_PORT")
	)

	var dsn = databaseUser + ":" + databasePass + "@tcp(" + databaseHost + ":" + databasePort + ")/" + databaseName

	if db, err := sql.Open("mysql", dsn); err != nil {
		return err
	} else {
		DB = db
	}

	return nil
}
