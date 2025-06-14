package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func ConnectDB() *sql.DB {
	_ = godotenv.Load() // ignore error if .env missing

	user := envOr("PG_USER", "app_user")
	pass := envOr("PG_PASS", "app_password")
	host := envOr("PG_HOST", "localhost")
	port := envOr("PG_PORT", "5432")
	dbName := envOr("PG_DB", "transfers")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, dbName)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("DB open: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("DB ping: %v", err)
	}
	return db
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
