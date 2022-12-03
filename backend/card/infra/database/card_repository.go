package database

import (
	"database/sql"

	"github.com/cidstein/super-brunfo/card/entity"
)

type CardRepository struct {
	Db *sql.DB
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{Db: db}
}

func (r *CardRepository) FindAll() ([]entity.Card, error) {
	rows, err := r.Db.Query("select id, name, attack, defense, intelligence, agility, resilience from card")

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
