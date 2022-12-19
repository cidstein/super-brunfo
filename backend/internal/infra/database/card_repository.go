package database

import (
	"context"

	"github.com/cidstein/super-brunfo/internal/model"
	"github.com/jackc/pgx/v5"
)

type CardRepositoryInterface interface {
	Save(ctx context.Context, card model.Card) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*model.Card, error)
	FindAll(ctx context.Context) ([]model.Card, error)
}

type CardRepository struct {
	Db *pgx.Conn
}

func NewCardRepository(db *pgx.Conn) *CardRepository {
	return &CardRepository{Db: db}
}

func (r *CardRepository) Save(ctx context.Context, card model.Card) error {
	_, err := r.Db.Exec(
		ctx,
		"INSERT INTO card (id, name, attack, defense, intelligence, agility, resilience, image_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		card.ID,
		card.Name,
		card.Attack,
		card.Defense,
		card.Intelligence,
		card.Agility,
		card.Resilience,
		card.ImageURL,
	)

	return err
}

func (r *CardRepository) Delete(ctx context.Context, id string) error {
	_, err := r.Db.Exec(
		ctx,
		"DELETE FROM card WHERE id = $1",
		id,
	)

	return err
}

func (r *CardRepository) FindByID(ctx context.Context, id string) (*model.Card, error) {
	var card model.Card

	err := r.Db.QueryRow(
		ctx,
		"SELECT id, name, attack, defense, intelligence, agility, resilience, image_url FROM card WHERE id = $1",
		id,
	).Scan(&card.ID, &card.Name, &card.Attack, &card.Defense, &card.Intelligence, &card.Agility, &card.Resilience, &card.ImageURL)

	return &card, err
}

func (r *CardRepository) FindAll(ctx context.Context) ([]model.Card, error) {
	rows, err := r.Db.Query(ctx, "select id, name, attack, defense, intelligence, agility, resilience, image_url from card")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []model.Card
	for rows.Next() {
		var card model.Card
		if err := rows.Scan(
			&card.ID,
			&card.Name,
			&card.Attack,
			&card.Defense,
			&card.Intelligence,
			&card.Agility,
			&card.Resilience,
			&card.ImageURL,
		); err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}
