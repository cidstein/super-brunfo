package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateNewMatch_ShouldReceiveError(t *testing.T) {
	match := Match{}
	assert.Error(t, match.IsValid(), "ID is required")
}

func TestGivenAnEmptyPlayerDeckId_WhenCreateNewMatch_ShouldReceiveError(t *testing.T) {
	match := Match{"1", "", "2", false, false}
	assert.Error(t, match.IsValid(), "player deck ID is required")
}

func TestGivenAnEmptyNpcDeckId_WhenCreateNewMatch_ShouldReceiveError(t *testing.T) {
	match := Match{"1", "1", "", false, false}
	assert.Error(t, match.IsValid(), "npc deck ID is required")
}

func TestGivenAValidMatch_WhenCreateNewMatch_ShouldReceiveNoError(t *testing.T) {
	match := Match{"1", "1", "2", false, false}
	assert.NoError(t, match.IsValid())
}
