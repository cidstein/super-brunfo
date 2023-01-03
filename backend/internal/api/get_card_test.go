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

	conn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := pgx.Connect(ctx, conn)
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	// Insert card
	_, err = db.Exec(
		ctx,
		`
			INSERT INTO card
				(id, name, attack, defense, intelligence, agility, resilience, flavour_text, image_url)
			VALUES
				($1, $2, $3, $4, $5, $6, $7, $8, $9)
		`,
		"141cbfd9-334e-4a4b-8559-3c6c4b50dc6c",
		"Test Card",
		10,
		10,
		10,
		10,
		10,
		"common",
		"https://example.com/image.png",
	)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/getcard?id=141cbfd9-334e-4a4b-8559-3c6c4b50dc6c", nil)
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

	// Delete card
	_, err = db.Exec(
		ctx,
		`
			DELETE FROM card
			WHERE id = $1
		`,
		"141cbfd9-334e-4a4b-8559-3c6c4b50dc6c",
	)
	if err != nil {
		t.Fatal(err)
	}
}
