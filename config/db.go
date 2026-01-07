package config

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(info map[string]string) (*pgxpool.Pool, error) {
	// Production
	// dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", info["USER"], info["PASSWORD"], info["URL"], info["DB_PORT"], info["DB_NAME"], info["SSL_MODE"])

	// Local (Peer)
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s", info["URL"], info["USER"], info["DB_NAME"], info["SSL_MODE"])

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
