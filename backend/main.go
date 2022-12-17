package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/cidstein/super-brunfo/internal/api"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()

	log.Print("Starting Super Brunfo! (backend) ...")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		// panic(err)
	}

	dbUrl := os.Getenv("DB_URL")
	log.Printf("DB_URL: %s", dbUrl)
	db, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Super Brunfo!"))
	})
	http.HandleFunc("/start", api.StartMatch(db))
	http.HandleFunc("/play", api.PlayGame(db))
	http.HandleFunc("/listcards", api.ListCards(db))

	err = http.ListenAndServe("0.0.0.0:"+os.Getenv("APP_PORT"), nil)
	if err != nil {
		fmt.Println("Error starting server")
		panic(err)
	}

}
