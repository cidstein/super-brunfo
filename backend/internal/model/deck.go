package model

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Deck struct {
	ID    string `json:"id"`
	Cards []Card `json:"cards"`
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

func (d *Deck) Split() (Deck, Deck, error) {
	if len(d.Cards) == 0 {
		return Deck{}, Deck{}, errors.New("deck is empty")
	}

	cut := len(d.Cards) / 2

	playerDeck := Deck{
		ID:    uuid.New().String(),
		Cards: d.Cards[:cut],
	}

	npcDeck := Deck{
		ID:    uuid.New().String(),
		Cards: d.Cards[cut:],
	}

	return playerDeck, npcDeck, nil
}

func (d *Deck) Draw() (Card, error) {
	if len(d.Cards) == 0 {
		return Card{}, errors.New("deck is empty")
	}

	card := d.Cards[0]
	d.Cards = d.Cards[1:]

	return card, nil
}

func (d *Deck) CheckIfEmpty() bool {
	return len(d.Cards) == 0
}
