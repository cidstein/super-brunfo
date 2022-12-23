package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"github.com/cidstein/super-brunfo/conf"
	"github.com/cidstein/super-brunfo/internal/api"
)

func main() {
	ctx := context.Background()

	log.Print("Starting Super Brunfo! (backend) ...")

	config, err := conf.Get()
	if err != nil {
		fmt.Println("Error getting config, %w", err)
		panic(err)
	}

	log.Printf("config.Version: %s", config.Version)
	log.Printf("config.DatabaseUrl: %s", config.DatabaseUrl)

	db, err := pgx.Connect(ctx, config.DatabaseUrl)
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Super Brunfo!"))
	})
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(config.Version))
	})
	http.HandleFunc("/start", api.StartMatch(db))
	http.HandleFunc("/play", api.PlayGame(db))
	http.HandleFunc("/playround", api.PlayRound(db))
	http.HandleFunc("/listcards", api.ListCards(db))
	http.HandleFunc("/listmatches", api.ListMatches(db))
	http.HandleFunc("/getcard", api.GetCard(db))
	http.HandleFunc("/getroundcards", api.GetRoundCards(db))

	address := fmt.Sprintf("%s:%s", config.AppHost, config.AppPort)
	log.Printf("Listening on %s", address)
	err = http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Error starting server")
		panic(err)
	}
}
