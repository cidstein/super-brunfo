package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateNewRound_ShouldReceiveError(t *testing.T) {
	round := Round{}
	assert.Error(t, round.IsValid(), "ID is required")
}

func TestGivenAnEmptyMatchId_WhenCreateNewRound_ShouldReceiveError(t *testing.T) {
	round := NewRound("1", "", "1", "1", 1, true, "1")
	assert.Error(t, round.IsValid(), "match ID is required")
}

func TestGivenAnEmptyPlayerCardId_WhenCreateNewRound_ShouldReceiveError(t *testing.T) {
	round := NewRound("1", "1", "", "1", 1, true, "1")
	assert.Error(t, round.IsValid(), "player card ID is required")
}

func TestGivenAnEmptyNpcCardId_WhenCreateNewRound_ShouldReceiveError(t *testing.T) {
	round := NewRound("1", "1", "1", "", 1, true, "1")
	assert.Error(t, round.IsValid(), "npc card ID is required")
}

func TestGivenAnEmptyAttribute_WhenCreateNewRound_ShouldReceiveError(t *testing.T) {
	round := NewRound("1", "1", "1", "1", 1, true, "")
	assert.Error(t, round.IsValid(), "attribute is required")
}

func TestGivenAValidRound_WhenCreateNewRound_ShouldNotReceiveError(t *testing.T) {
	round := NewRound("1", "1", "1", "1", 1, true, "1")
	assert.NoError(t, round.IsValid())
}
