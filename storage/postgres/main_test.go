package postgres_test

import (
	"fmt"
	"github.com/MuhammadyusufAdhamov/medium_user_service/config"
	"github.com/MuhammadyusufAdhamov/medium_user_service/storage"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"testing"
)

var (
	strg storage.StorageI
)

func TestMain(m *testing.M) {
	cfg := config.Load("./../..")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	strg = storage.NewStoragePg(db)
	os.Exit(m.Run())
}
