package db

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DSN string `toml:"dsn" validate:"required" env:"DB_DSN,overwrite"`
}

type Client struct {
	db *sql.DB
}

func New(c Config) (*Client, error) {
	// initialise db and return client struct
	return nil, nil
}

func (c *Client) Close() error {
	// close db instance; to be called from main when process shuts down
	return nil
}

func (c *Client) UpdateFeedback(ctx context.Context, feedback *Feedback) error {
	return nil
}
