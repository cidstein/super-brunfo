package database

import (
	"github.com/cidstein/super-brunfo/card/entity"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

type DeckRepository struct {
	Db *pgx.Conn
}

func NewDeckRepository(db *pgx.Conn) *DeckRepository {
	return &DeckRepository{Db: db}
}

func (r *DeckRepository) Save() (*entity.Deck, error) {
	rows, err := r.Db.Query("SELECT id from cards")
	if err != nil {
		return nil, err
	}

	var cards []entity.Card
	for rows.Next() {
		var card entity.Card
		err = rows.Scan(&card.ID)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	id := uuid.New().String()
	deck := entity.NewDeck(id, cards)

	_, err = r.Db.Exec(
		"INSERT INTO deck (id) VALUES ($1)", deck.ID,
	)
	if err != nil {
		return nil, err
	}

	for _, card := range deck.Cards {
		_, err = r.Db.Exec(
			"INSERT INTO deck_cards (deck_id, card_id) VALUES ($1, $2)", deck.ID, card.ID,
		)
		if err != nil {
			return nil, err
		}
	}

	return &deck, nil
}

func (r *DeckRepository) FindByID(id string) (entity.Deck, error) {
	var deck entity.Deck

	err := r.Db.QueryRow(
		"SELECT id FROM deck WHERE id = $1",
		id,
	).Scan(&deck.ID)
	if err != nil {
		return deck, err
	}

	rows, err := r.Db.Query(
		"SELECT card_id FROM deck_cards WHERE deck_id = $1",
		id,
	)
	if err != nil {
		return deck, err
	}

	var cards []entity.Card
	for rows.Next() {
		var card entity.Card
		err = rows.Scan(&card.ID)
		if err != nil {
			return deck, err
		}
		cards = append(cards, card)
	}

	deck.Cards = cards

	return deck, nil
}
