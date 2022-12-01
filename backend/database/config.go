package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func Open() *pgx.Conn {
	ctx := context.Background()

	db, err := pgx.Connect(ctx, "postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(ctx)

	err = db.Ping(ctx)
	if err != nil {
		fmt.Println("Error pinging database: ", err)
		panic(err)
	}

	fmt.Println("Connected to database")

	return db
}
