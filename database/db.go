package database

import (
	"awesomeProject/users"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type database struct {
	pool *pgxpool.Pool
}

// New - construct DB entity
func New(ctx context.Context, url string) (*database, error) {
	pool, err := pgxpool.Connect(ctx, url)
	if err != nil {
		return &database{}, err
	}

	return &database{pool: pool}, nil
}

// CreateSchema - creates table if not exist users.
func (db *database) CreateSchema(ctx context.Context) error {
	createSchemeQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id BYTEA PRIMARY KEY NOT NULL, 
			email VARCHAR NOT NULL,
			name VARCHAR NOT NULL,
			password BYTEA NOT NULL,
			created_at 	TIMESTAMP WITH TIME ZONE NOT NULL
		);`

	_, err := db.pool.Exec(ctx, createSchemeQuery)
	if err != nil {
		return err
	}

	return nil
}

// Close closes database connection.
func (db *database) Close() {
	db.pool.Close()
}

func (db *database) Users() users.DB {
	return &userDB{pool: db.pool}
}
