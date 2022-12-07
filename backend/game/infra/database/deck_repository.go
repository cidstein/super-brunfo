package database

import (
	"context"

	"github.com/cidstein/super-brunfo/game/entity"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type DeckRepository struct {
	Db *pgx.Conn
}

func NewDeckRepository(db *pgx.Conn) *DeckRepository {
	return &DeckRepository{Db: db}
}

func (r *DeckRepository) Save(ctx context.Context) (*entity.Deck, error) {
	rows, err := r.Db.Query(ctx, "SELECT id from cards")
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
		ctx,
		"INSERT INTO deck (id) VALUES ($1)",
		deck.ID,
	)
	if err != nil {
		return nil, err
	}

	for _, card := range deck.Cards {
		_, err = r.Db.Exec(
			ctx,
			"INSERT INTO deck_cards (deck_id, card_id) VALUES ($1, $2)",
			deck.ID,
			card.ID,
		)
		if err != nil {
			return nil, err
		}
	}

	return &deck, nil
}

func (r *DeckRepository) Delete(ctx context.Context, id string) error {
	_, err := r.Db.Exec(
		ctx,
		"DELETE FROM deck_cards WHERE deck_id = $1",
		id,
	)
	if err != nil {
		return err
	}

	_, err = r.Db.Exec(
		ctx,
		"DELETE FROM deck WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *DeckRepository) FindByID(ctx context.Context, id string) (entity.Deck, error) {
	var deck entity.Deck

	err := r.Db.QueryRow(
		ctx,
		"SELECT id FROM deck WHERE id = $1",
		id,
	).Scan(&deck.ID)
	if err != nil {
		return deck, err
	}

	rows, err := r.Db.Query(
		ctx,
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
