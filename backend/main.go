package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"github.com/cidstein/super-brunfo/conf"
	"github.com/cidstein/super-brunfo/internal/api"
	"github.com/cidstein/super-brunfo/internal/middleware"
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

	router := gin.Default()

	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", []byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

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

	router.POST("/signup", api.SignUp(db))
	router.POST("/signin", api.SignIn(db))
	router.POST("/signout", api.SignOut(db))

	auth := router.Group("/auth")
	auth.Use(middleware.Auth())
	{
		auth.GET("/", api.Home())
		auth.GET("/version", api.Version(config.Version))
		auth.POST("/start", api.StartMatch(db))
		auth.GET("/loadround", api.LoadRound(db))
		auth.PUT("/playround", api.PlayRound(db))
		auth.GET("/listcards", api.ListCards(db))
		auth.GET("/listmatches", api.ListMatches(db))
	}

	router.Run()
}
