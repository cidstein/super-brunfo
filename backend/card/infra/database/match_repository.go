package database

import (
	"database/sql"

	"github.com/cidstein/super-brunfo/card/entity"
)

type MatchRepository struct {
	Db *sql.DB
}

func NewMatchRepository(db *sql.DB) *MatchRepository {
	return &MatchRepository{Db: db}
}

func (r *MatchRepository) Save(match entity.Match) error {
	_, err := r.Db.Exec(
		"insert into match (id, deck1, deck2, winner) values (?, ?, ?, ?)",
		match.ID,
		match.Deck1,
		match.Deck2,
		match.Winner,
	)

	return err
}
