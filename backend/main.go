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
		fmt.Println("Error getting config")
		panic(err)
	}

	log.Printf("DB_URL: %s", config.DatabaseUrl)
	db, err := pgx.Connect(ctx, config.DatabaseUrl)
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	log.Printf("config.Version: %s", config.Version)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Super Brunfo!"))
	})
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(config.Version))
	})
	http.HandleFunc("/start", api.StartMatch(db))
	http.HandleFunc("/play", api.PlayGame(db))
	http.HandleFunc("/listcards", api.ListCards(db))

	address := fmt.Sprintf("%s:%s", config.AppHost, config.AppPort)
	log.Printf("Listening on %s", address)
	err = http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Error starting server")
		panic(err)
	}
}
