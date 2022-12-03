package entity

import (
	"errors"
	"math/rand"
	"time"
)

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

	if len(d.Cards)%2 != 0 {
		return errors.New("deck must have an even number of cards")
	}

	return nil
}

func (d *Deck) Shuffle() (Deck, error) {
	if err := d.IsValid(); err != nil {
		return Deck{}, err
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })

	return *d, nil
}

func (d *Deck) Cut() (Deck, Deck, error) {
	if len(d.Cards) == 0 {
		return Deck{}, Deck{}, errors.New("deck is empty")
	}

	cut := len(d.Cards) / 2

	deck1 := Deck{
		ID:    d.ID,
		Cards: d.Cards[:cut],
	}

	deck2 := Deck{
		ID:    d.ID,
		Cards: d.Cards[cut:],
	}

	return deck1, deck2, nil
}
