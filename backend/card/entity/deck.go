package entity

import (
	"errors"
)

type Deck struct {
	ID string
}

type CardInDeck struct {
	DeckID string
	CardID string
}

func NewDeck(id string) Deck {
	return Deck{
		ID: id,
	}
}

func NewCardInDeck(deckID string, cardID string) CardInDeck {
	return CardInDeck{
		DeckID: deckID,
		CardID: cardID,
	}
}

func NewCardInDeckBatch(deckID string, cardIDs []string) []CardInDeck {
	var cardsInDeck []CardInDeck

	for _, cardID := range cardIDs {
		cardsInDeck = append(cardsInDeck, NewCardInDeck(deckID, cardID))
		NewCardInDeck(deckID, cardID)
	}

	return cardsInDeck
}

func (d *Deck) IsValid() error {
	if d.ID == "" {
		return errors.New("ID is required")
	}

	return nil
}

func (c *CardInDeck) IsValid() error {
	if c.DeckID == "" {
		return errors.New("deck ID is required")
	}

	if c.CardID == "" {
		return errors.New("card ID is required")
	}

	return nil
}

///

// func (d *Deck) Shuffle() (DeckInCards, error) {
// 	if len(d.Cards) == 0 {
// 		return Deck{}, errors.New("deck is empty")
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })

// 	return *d, nil
// }

// func (d *Deck) Cut() (Deck, Deck, error) {
// 	if len(d.Cards) == 0 {
// 		return Deck{}, Deck{}, errors.New("deck is empty")
// 	}

// 	cut := len(d.Cards) / 2

// 	deck1 := Deck{
// 		ID:    d.ID,
// 		Cards: d.Cards[:cut],
// 	}

// 	deck2 := Deck{
// 		ID:    d.ID,
// 		Cards: d.Cards[cut:],
// 	}

// 	return deck1, deck2, nil
// }
