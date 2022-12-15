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
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	dbUrl := os.Getenv("DB_URL")
	db, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	http.HandleFunc("/start", handlers.StartMatch(db))
	http.HandleFunc("/play", handlers.PlayGame(db))
	http.HandleFunc("/listcards", handlers.ListCards(db))

	http.ListenAndServe(":8080", nil)
}
