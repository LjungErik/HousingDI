package sql

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

// Config Config that contains connection string to
// sql database
type Config struct {
	ConnString string
}

// Client client containing the connection pool
type Client struct {
	pool *pgxpool.Pool
}

// NewClient Creates a new SQL client
// based on the provided config
func NewClient(conf *Config) *Client {
	pool, err := pgxpool.Connect(context.Background(), conf.ConnString)
	if err != nil {
		log.Warn("Failed to setup connection to sql server")
		panic(err)
	}

	// ping database to check connection works
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Error("Failed to acquire connection to sql server")
		panic(err)
	}

	defer conn.Release()

	err = conn.Conn().Ping(context.Background())
	if err != nil {
		log.Error("Failed to ping sql server")
		panic(err)
	}

	return &Client{
		pool: pool,
	}
}
