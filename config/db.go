package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func InitDB(info map[string]string) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", info["USER"], info["PASSWORD"], info["URL"], info["PORT"], info["DB_NAME"])
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
}
