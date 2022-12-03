package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{}
	assert.Error(t, card.IsValid(), "ID is required")
}

func TestGivenAnEmptyName_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{"1", "", 0, 0, 0, 0, 0}
	assert.Error(t, card.IsValid(), "name is required")
}

func TestGivenAnEmptyAttack_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{"1", "name", -1, 0, 0, 0, 0}
	assert.Error(t, card.IsValid(), "attack must be between 0 and 100")
}

func TestGivenAnEmptyDefence_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{"1", "name", 0, -1, 0, 0, 0}
	assert.Error(t, card.IsValid(), "defense must be between 0 and 100")
}

func TestGivenAnEmptyIntelligence_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{"1", "name", 0, 0, -1, 0, 0}
	assert.Error(t, card.IsValid(), "intelligence must be between 0 and 100")
}

func TestGivenAnEmptyAgility_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{"1", "name", 0, 0, 0, -1, 0}
	assert.Error(t, card.IsValid(), "agility must be between 0 and 100")
}

func TestGivenAnEmptyResilience_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{"1", "name", 0, 0, 0, 0, -1}
	assert.Error(t, card.IsValid(), "resilience must be between 0 and 100")
}

func TestGivenAValidCard_WhenCreateNewCard_ShouldReceiveNoError(t *testing.T) {
	card := Card{"1", "name", 0, 0, 0, 0, 0}
	assert.NoError(t, card.IsValid())
}

func TestGivenAInvalidCard_WhenCallCombat_ShouldReceiveAnError(t *testing.T) {
	card1 := NewCard("1", "name1", 0, 0, 0, 0, 0)
	card2 := NewCard("2", "name2", -10, 0, 0, 0, 0)

	_, err := card1.Combat(&card2, "attack")
	assert.Error(t, err)
}

func TestGivenAInvalidAttribute_WhenCallCombat_ShouldReceiveAnError(t *testing.T) {
	card1 := NewCard("1", "name1", 0, 0, 0, 0, 0)
	card2 := NewCard("2", "name2", 0, 0, 0, 0, 0)

	_, err := card1.Combat(&card2, "power")
	assert.Error(t, err)
}

func TestGivenAValidCard_WhenCallCombat_ShouldReceiveCard1AsWinner(t *testing.T) {
	card1 := NewCard("1", "name1", 10, 0, 0, 0, 0)
	card2 := NewCard("2", "name2", 0, 0, 0, 0, 0)

	winner, err := card1.Combat(&card2, "attack")
	assert.NoError(t, err)
	assert.Equal(t, true, winner)
}
func TestGivenAValidCard_WhenCallCombat_ShouldReceiveCard2AsWinner(t *testing.T) {
	card1 := NewCard("1", "name1", 0, 0, 0, 0, 0)
	card2 := NewCard("2", "name2", 10, 0, 0, 0, 0)

	winner, err := card1.Combat(&card2, "attack")
	assert.NoError(t, err)
	assert.Equal(t, false, winner)
}
