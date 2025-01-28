package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //  `sql.Open`에서 `postgres` driver를 사용하기 위해 import
)

// Config represents the configuration for the Postgres database.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// New creates a new SQLX database connection.
func New(cfg Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Connected to PostgreSQL database!")

	return db, nil
}
