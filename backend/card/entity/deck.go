package entity

import "errors"

type Deck struct {
	ID    string
	Cards []Card
}

func NewDeck(id string, cards []Card) Deck {
	return Deck{
		ID:    id,
		Cards: cards,
	}
}

func (d *Deck) IsValid() error {
	if d.ID == "" {
		return errors.New("ID is required")
	}

	if len(d.Cards) == 0 {
		return errors.New("deck must have at least one card")
	}

	return nil
}

func (d *Deck) PickCard() (Card, error) {
	if len(d.Cards) == 0 {
		return Card{}, errors.New("deck is empty")
	}

	card := d.Cards[0]
	d.Cards = d.Cards[1:]

	return card, nil
}
