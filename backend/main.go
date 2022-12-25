package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
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
		panic(err)
	}

	/* Gin router */
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "https://github.com"},
		AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS", "PATCH"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Methods",
			"Access-Control-Allow-Credentials",
			"Access-Control-Expose-Headers",
			"Access-Control-Max-Age",
			"Access-Control-Request-Headers",
			"Access-Control-Request-Method",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Content-Type",
			"Authorization",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Methods",
			"Access-Control-Allow-Credentials",
			"Access-Control-Expose-Headers",
			"Access-Control-Max-Age",
			"Access-Control-Request-Headers",
			"Access-Control-Request-Method",
		},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/", api.Home())
	router.GET("/version", api.Version(config.Version))
	router.GET("/start", api.StartMatch(db))
	router.GET("/loadround", api.LoadRound(db))
	router.PUT("/playround", api.PlayRound(db))
	router.GET("/listcards", api.ListCards(db))
	router.GET("/listmatches", api.ListMatches(db))
	// http.HandleFunc("/getcard", api.GetCard(db))
	// http.HandleFunc("/getroundcards", api.GetRoundCards(db))

	router.Run()
}
