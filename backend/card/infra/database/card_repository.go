package database

import (
	"context"

	"github.com/cidstein/super-brunfo/card/entity"
	"github.com/jackc/pgx/v5"
)

type CardRepository struct {
	ctx context.Context
	Db  *pgx.Conn
}

func NewCardRepository(db *pgx.Conn) *CardRepository {
	return &CardRepository{Db: db}
}

func (r *CardRepository) Save(card entity.Card) error {
	_, err := r.Db.Exec(
		context.Background(),
		"INSERT INTO card (name, attack, defense, intelligence, agility, resilience) VALUES ($1, $2, $3, $4, $5, $6)",
		card.Name,
		card.Attack,
		card.Defense,
		card.Intelligence,
		card.Agility,
		card.Resilience,
	)

	return err
}

func (r *CardRepository) FindByID(id string) (entity.Card, error) {
	var card entity.Card

	err := r.Db.QueryRow(
		context.Background(),
		"SELECT id, name, attack, defense, intelligence, agility, resilience FROM card WHERE id = $1",
		id,
	).Scan(&card.ID, &card.Name, &card.Attack, &card.Defense, &card.Intelligence, &card.Agility, &card.Resilience)

	return card, err
}

func (r *CardRepository) FindAll() ([]entity.Card, error) {
	rows, err := r.Db.Query(r.ctx, "select id, name, attack, defense, intelligence, agility, resilience from card")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []entity.Card
	for rows.Next() {
		var card entity.Card
		if err := rows.Scan(
			&card.ID,
			&card.Name,
			&card.Attack,
			&card.Defense,
			&card.Intelligence,
			&card.Agility,
			&card.Resilience,
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
