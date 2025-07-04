package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
)

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func ConnectDB() *gorm.DB {
	// Load .env.test during testing
	if os.Getenv("GO_ENV") == "test" {
		log.Printf("loading .env.test")
		_ = godotenv.Load("../.env.test")
	} else {
		_ = godotenv.Load()
	}

	user := envOr("PG_USER", "postgres")
	pass := envOr("PG_PASS", "password")
	host := envOr("PG_HOST", "localhost")
	port := envOr("PG_PORT", "5432")
	dbName := envOr("PG_DB", "transfers")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, pass, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return db
}

