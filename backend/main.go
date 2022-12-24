package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
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
		// panic(err)
	}

	/* Gin router */
	router := gin.Default()

	// router.Use(cors.Middleware(cors.Config{
	// 	Origins:         "*",
	// 	Methods:         "GET, PUT, POST, DELETE",
	// 	RequestHeaders:  "Origin, Authorization, Content-Type",
	// 	ExposedHeaders:  "",
	// 	MaxAge:          50 * time.Second,
	// 	Credentials:     true,
	// 	ValidateHeaders: false,
	// }))

	router.GET("/", api.Home())
	router.GET("/version", api.Version(config.Version))
	// http.HandleFunc("/start", api.StartMatch(db))
	router.GET("/loadround", api.LoadRound(db))
	router.PUT("/playround", api.PlayRound(db))
	router.GET("/listcards", api.ListCards(db))
	// http.HandleFunc("/listmatches", api.ListMatches(db))
	// http.HandleFunc("/getcard", api.GetCard(db))
	// http.HandleFunc("/getroundcards", api.GetRoundCards(db))

	router.Run()
}
