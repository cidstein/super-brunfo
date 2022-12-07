package main

import (
	"context"
	"net/http"

	"github.com/cidstein/super-brunfo/game/handlers"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/start", handlers.StartMatch(db))
	http.HandleFunc("/play", handlers.PlayGame(db))

	http.ListenAndServe(":8080", nil)
}
