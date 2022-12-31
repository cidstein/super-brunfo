package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{}
	assert.Error(t, card.IsValid(), "ID is required")
}

func TestGivenAnEmptyName_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{"1", "", 0, 0, 0, 0, 0, "", ""}
	assert.Error(t, card.IsValid(), "name is required")
}

func TestGivenAnEmptyAttack_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{"1", "name", -1, 0, 0, 0, 0, "", ""}
	assert.Error(t, card.IsValid(), "attack must be between 0 and 100")
}

func TestGivenAnEmptyDefence_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{"1", "name", 0, -1, 0, 0, 0, "", ""}
	assert.Error(t, card.IsValid(), "defense must be between 0 and 100")
}

func TestGivenAnEmptyIntelligence_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{"1", "name", 0, 0, -1, 0, 0, "", ""}
	assert.Error(t, card.IsValid(), "intelligence must be between 0 and 100")
}

func TestGivenAnEmptyAgility_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{"1", "name", 0, 0, 0, -1, 0, "", ""}
	assert.Error(t, card.IsValid(), "agility must be between 0 and 100")
}

func TestGivenAnEmptyResilience_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{"1", "name", 0, 0, 0, 0, -1, "", ""}
	assert.Error(t, card.IsValid(), "resilience must be between 0 and 100")
}

func TestGivenAnEmptyFlavourText_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{"1", "name", 0, 0, 0, 0, 0, "", ""}
	assert.Error(t, card.IsValid(), "flavour_text is required")
}

func TestGivenAnEmptyImageURL_WhenCreateNewCard_ShouldReceiveError(t *testing.T) {
	card := Card{"1", "name", 0, 0, 0, 0, 0, "flavour_text", ""}
	assert.Error(t, card.IsValid(), "image_url is required")
}

func TestGivenAValidCard_WhenCreateNewCard_ShouldReceiveNoError(t *testing.T) {
	card := Card{"1", "name", 0, 0, 0, 0, 0, "flavour_text", "https://1drv.ms/u/s!Aq4ssY1EMmRWh9o6Ax8aMkel7eetng?e=s1TMI2"}
	assert.NoError(t, card.IsValid())
}

func TestGivenAInvalidPlayerCard_WhenCallCombat_ShouldReceiveAnError(t *testing.T) {
	card1 := NewCard("1", "name1", -10, 0, 0, 0, 0, "", "")
	card2 := NewCard("2", "name2", 0, 0, 0, 0, 0, "", "")

	_, err := card1.Combat(&card2, "attack")
	assert.Error(t, err)
}

func TestGivenAInvalidNpcCard_WhenCallCombat_ShouldReceiveAnError(t *testing.T) {
	card1 := NewCard("1", "name1", 0, 0, 0, 0, 0, "", "")
	card2 := NewCard("2", "name2", -10, 0, 0, 0, 0, "", "")

	_, err := card1.Combat(&card2, "attack")
	assert.Error(t, err)
}

func TestGivenAInvalidAttribute_WhenCallCombat_ShouldReceiveAnError(t *testing.T) {
	card1 := NewCard("1", "name1", 0, 0, 0, 0, 0, "", "")
	card2 := NewCard("2", "name2", 0, 0, 0, 0, 0, "", "")

	_, err := card1.Combat(&card2, "power")
	assert.Error(t, err)
}

func TestGivenAValidCardWithHigherAttack_WhenCallCombat_ShouldReceiveCard1AsWinner(t *testing.T) {
	card1 := NewCard("1", "name1", 10, 0, 0, 0, 0, "a", "https://i.imgur.com/PsF78Ls.jpg")
	card2 := NewCard("2", "name2", 0, 0, 0, 0, 0, "a", "https://i.imgur.com/PsF78Ls.jpg")

	winner, err := card1.Combat(&card2, "attack")
	assert.NoError(t, err)
	assert.Equal(t, true, winner)
}

func TestGivenAValidCardWithHigherDefense_WhenCallCombat_ShouldReceiveCard1AsWinner(t *testing.T) {
	card1 := NewCard("1", "name1", 0, 10, 0, 0, 0, "a", "https://i.imgur.com/PsF78Ls.jpg")
	card2 := NewCard("2", "name2", 0, 0, 0, 0, 0, "a", "https://i.imgur.com/PsF78Ls.jpg")

	winner, err := card1.Combat(&card2, "defense")
	assert.NoError(t, err)
	assert.Equal(t, true, winner)
}

func TestGivenAValidCardWithHigherIntelligence_WhenCallCombat_ShouldReceiveCard1AsWinner(t *testing.T) {
	card1 := NewCard("1", "name1", 0, 0, 10, 0, 0, "a", "https://i.imgur.com/PsF78Ls.jpg")
	card2 := NewCard("2", "name2", 0, 0, 0, 0, 0, "a", "https://i.imgur.com/PsF78Ls.jpg")

	winner, err := card1.Combat(&card2, "intelligence")
	assert.NoError(t, err)
	assert.Equal(t, true, winner)
}

func TestGivenAValidCardWithHigherAgility_WhenCallCombat_ShouldReceiveCard1AsWinner(t *testing.T) {
	card1 := NewCard("1", "name1", 0, 0, 0, 10, 0, "a", "https://i.imgur.com/PsF78Ls.jpg")
	card2 := NewCard("2", "name2", 0, 0, 0, 0, 0, "a", "https://i.imgur.com/PsF78Ls.jpg")

	winner, err := card1.Combat(&card2, "agility")
	assert.NoError(t, err)
	assert.Equal(t, true, winner)
}

func TestGivenAValidCardWithHigherResilience_WhenCallCombat_ShouldReceiveCard1AsWinner(t *testing.T) {
	card1 := NewCard("1", "name1", 0, 0, 0, 0, 10, "a", "https://i.imgur.com/PsF78Ls.jpg")
	card2 := NewCard("2", "name2", 0, 0, 0, 0, 0, "a", "https://i.imgur.com/PsF78Ls.jpg")

	winner, err := card1.Combat(&card2, "resilience")
	assert.NoError(t, err)
	assert.Equal(t, true, winner)
}

func TestGivenAValidCardWithInvalidAttribute_WhenCallCombat_ShouldReceiveCard1AsWinner(t *testing.T) {
	card1 := NewCard("1", "name1", 0, 0, 0, 0, 0, "a", "https://i.imgur.com/PsF78Ls.jpg")
	card2 := NewCard("2", "name2", 0, 0, 0, 0, 0, "a", "https://i.imgur.com/PsF78Ls.jpg")

	_, err := card1.Combat(&card2, "invalid")
	assert.Error(t, err)
}

func TestGivenAnInvalidCard_WhenCallCombat_ShouldReceiveError(t *testing.T) {
	card1 := NewCard("1", "name1", 0, 0, 0, 0, 0, "a", "https://i.imgur.com/PsF78Ls.jpg")
	card2 := Card{}

	_, err := card1.Combat(&card2, "attack")
	assert.Error(t, err)
}
