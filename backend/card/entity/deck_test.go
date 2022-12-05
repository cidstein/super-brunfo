package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateNewDeck_ShouldReceiveError(t *testing.T) {
	deck := Deck{}
	assert.Error(t, deck.IsValid(), "ID is required")
}

func TestGivenAnEmptyCards_WhenCreateNewDeck_ShouldReceiveError(t *testing.T) {
	deck := Deck{"1", []Card{}}
	assert.Error(t, deck.IsValid(), "deck must have at least one card")
}

func TestGivenAnOddNumberOfCards_WhenCreateNewDeck_ShouldReceiveError(t *testing.T) {
	deck := Deck{"1", []Card{{"1", "name", 0, 0, 0, 0, 0}}}
	assert.Error(t, deck.IsValid(), "deck must have an even number of cards")
}

func TestGivenAValidDeck_WhenCreateNewDeck_ShouldReceiveNoError(t *testing.T) {
	deck := Deck{"1", []Card{{"1", "name", 0, 0, 0, 0, 0}, {"2", "name", 0, 0, 0, 0, 0}}}
	assert.NoError(t, deck.IsValid())
}

func TestGivenAnEmptyDeck_WhenShuffleDeck_ShouldReceiveError(t *testing.T) {
	deck := Deck{}
	_, err := deck.Shuffle()
	assert.Error(t, err, "ID is required")
}

func TestGivenAValidDeck_WhenShuffleDeck_ShouldReceiveNoError(t *testing.T) {
	deck := Deck{"1", []Card{{"1", "name", 0, 0, 0, 0, 0}, {"2", "name", 0, 0, 0, 0, 0}}}
	_, err := deck.Shuffle()
	assert.NoError(t, err)
}

func TestGivenAnEmptyDeck_WhenSplitDeck_ShouldReceiveError(t *testing.T) {
	deck := Deck{}
	_, _, err := deck.Split()
	assert.Error(t, err, "deck is empty")
}

func TestGivenAValidDeck_WhenSplitDeck_ShouldReceiveNoError(t *testing.T) {
	deck := Deck{"1", []Card{{"1", "name", 0, 0, 0, 0, 0}, {"2", "name", 0, 0, 0, 0, 0}}}
	_, _, err := deck.Split()
	assert.NoError(t, err)
}

func TestGivenAValidDeck_WhenSplitDeck_ShouldReceiveTwoDecks(t *testing.T) {
	deck := Deck{"1", []Card{{"1", "name", 0, 0, 0, 0, 0}, {"2", "name", 0, 0, 0, 0, 0}}}
	playerDeck, npcDeck, _ := deck.Split()
	assert.Equal(t, 1, len(playerDeck.Cards))
	assert.Equal(t, 1, len(npcDeck.Cards))
}

func TestGivenAValidDeck_WhenSplitDeck_ShouldReceiveTwoDecksWithTheSameCards(t *testing.T) {
	deck := Deck{"1", []Card{{"1", "name", 0, 0, 0, 0, 0}, {"2", "name", 0, 0, 0, 0, 0}}}
	playerDeck, npcDeck, _ := deck.Split()
	assert.Equal(t, deck.Cards[0], playerDeck.Cards[0])
	assert.Equal(t, deck.Cards[1], npcDeck.Cards[0])
}

func TestGivenAValidDeck_WhenSplitDeck_ShouldReceiveTwoDecksWithDifferentIds(t *testing.T) {
	deck := Deck{"1", []Card{{"1", "name", 0, 0, 0, 0, 0}, {"2", "name", 0, 0, 0, 0, 0}}}
	playerDeck, npcDeck, _ := deck.Split()
	assert.NotEqual(t, playerDeck.ID, npcDeck.ID)
}

func TestGivenAnEmptyDeck_WhenDrawCard_ShouldReceiveError(t *testing.T) {
	deck := Deck{}
	_, err := deck.Draw()
	assert.Error(t, err, "deck is empty")
}

func TestGivenAValidDeck_WhenDrawCard_ShouldReceiveNoError(t *testing.T) {
	deck := Deck{"1", []Card{{"1", "name", 0, 0, 0, 0, 0}, {"2", "name", 0, 0, 0, 0, 0}}}
	_, err := deck.Draw()
	assert.NoError(t, err)
}

func TestGivenAValidDeck_WhenDrawCard_ShouldReceiveTheFirstCard(t *testing.T) {
	deck := Deck{"1", []Card{{"1", "name", 0, 0, 0, 0, 0}, {"2", "name", 0, 0, 0, 0, 0}}}
	firstCard := Card{"1", "name", 0, 0, 0, 0, 0}
	card, _ := deck.Draw()
	assert.Equal(t, firstCard, card)
}

func TestGivenAValidDeck_WhenDrawCard_ShouldReceiveTheDeckWithoutTheFirstCard(t *testing.T) {
	deck := Deck{"1", []Card{{"1", "name", 0, 0, 0, 0, 0}, {"2", "name", 0, 0, 0, 0, 0}}}
	deck.Draw()
	assert.Equal(t, 1, len(deck.Cards))
}

func TestGivenAnEmptyDeck_WhenCheckIfEmpty_ShouldReceiveTrue(t *testing.T) {
	deck := Deck{}
	assert.True(t, deck.CheckIfEmpty())
}

func TestGivenAValidDeck_WhenCheckIfEmpty_ShouldReceiveFalse(t *testing.T) {
	deck := Deck{"1", []Card{{"1", "name", 0, 0, 0, 0, 0}, {"2", "name", 0, 0, 0, 0, 0}}}
	assert.False(t, deck.CheckIfEmpty())
}
