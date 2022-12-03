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
