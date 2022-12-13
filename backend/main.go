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

	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		fmt.Println("Error loading DB_USER")
	}

	dbPass, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		fmt.Println("Error loading DB_PASSWORD")
	}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		fmt.Println("Error loading DB_HOST")
	}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		fmt.Println("Error loading DB_PORT")
	}

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		fmt.Println("Error loading DB_NAME")
	}

	conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
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
