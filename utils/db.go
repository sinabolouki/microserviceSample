package utils

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func InitDB() *sql.DB {
	// fallback is not for production
	connStr := getEnv("POSTGRES_CONN",
		"postgres://postgres:postgres@localhost:5432/microservice_sample?sslmode=disable") // e.g., postgres://user:pass@host:port/dbname?sslmode=disable
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("DB unreachable: %v", err)
	}
	return db
}
