package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/cidstein/super-brunfo/game/handlers"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	ctx := context.Background()

	conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println(conn)

	db, err := pgx.Connect(ctx, conn)
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	http.HandleFunc("/start", handlers.StartMatch(db))
	http.HandleFunc("/play", handlers.PlayGame(db))

	http.ListenAndServe(":8080", nil)
}
