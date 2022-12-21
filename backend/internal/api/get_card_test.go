package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func TestGetCardHandler(t *testing.T) {
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

	req, err := http.NewRequest("GET", "/getcard?id=9f240160-ed87-44db-b9d9-07a4ea180c28", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCard(db))

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
