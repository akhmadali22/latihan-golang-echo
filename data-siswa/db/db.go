package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	connStr, err := loadPostgresConfig()
	if err != nil {
		return
	}

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic("connectionString Error")
	}

	err = db.Ping()

	if err != nil {
		panic("DSN Error" + connStr)
	}
}

func loadPostgresConfig() (string, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		// os.Getenv("DB_USER"),
		"postgres",
		// os.Getenv("DB_PASSWORD"),
		"rahasia",
		// os.Getenv("DB_HOST"),
		"192.168.1.88",
		// os.Getenv("DB_PORT"),
		"5430",
		// os.Getenv("DB_DATABASE"),
		"belajar",
	)
	return connStr, nil
}

func CreateCon() *sql.DB {
	return db
}
