package database

import "database/sql"

type DeckRepository struct {
	Db *sql.DB
}

func NewDeckRepository(db *sql.DB) *DeckRepository {
	return &DeckRepository{Db: db}
}

// func (r *DeckRepository) Save() ([]entity.Deck, error) {
