package repository

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/cidstein/super-brunfo/internal/model"
)

type DeckRepositoryInterface interface {
	Save(ctx context.Context, deck model.Deck) (*model.Deck, error)
	Delete(ctx context.Context, id string) error
	DrawCard(ctx context.Context, deckID, cardID string) error
	FindByID(ctx context.Context, id string) (*model.Deck, error)
}

type DeckRepository struct {
	Db *pgx.Conn
}

func NewDeckRepository(db *pgx.Conn) *DeckRepository {
	return &DeckRepository{Db: db}
}

func (r *DeckRepository) Save(ctx context.Context, deck model.Deck) (*model.Deck, error) {
	_, err := r.Db.Exec(
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

func (r *DeckRepository) DrawCard(ctx context.Context, deckID, cardID string) error {
	_, err := r.Db.Exec(
		ctx,
		"DELETE FROM deck_cards WHERE deck_id = $1 AND card_id = $2",
		deckID,
		cardID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *DeckRepository) FindByID(ctx context.Context, id string) (*model.Deck, error) {
	var deck model.Deck

	err := r.Db.QueryRow(
		ctx,
		"SELECT id FROM deck WHERE id = $1",
		id,
	).Scan(&deck.ID)
	if err != nil {
		return nil, err
	}

	rows, err := r.Db.Query(
		ctx,
		"SELECT card_id FROM deck_cards WHERE deck_id = $1",
		id,
	)
	if err != nil {
		return nil, err
	}

	var cards []model.Card
	for rows.Next() {
		var card model.Card
		err = rows.Scan(&card.ID)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	deck.Cards = cards

	return &deck, nil
}
