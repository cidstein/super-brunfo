package entity

import "errors"

type Match struct {
	ID           string
	PlayerDeckID string
	NpcDeckID    string
	Victory      bool
	Finished     bool
}

func NewMatch(id string, deckPlayerID string, deckComID string, victory, finished bool) Match {
	return Match{
		ID:           id,
		PlayerDeckID: deckPlayerID,
		NpcDeckID:    deckComID,
		Victory:      victory,
		Finished:     finished,
	}
}

func (m Match) IsValid() error {
	if m.ID == "" {
		return errors.New("ID is required")
	}

	if m.PlayerDeckID == "" {
		return errors.New("player deck ID is required")
	}

	if m.NpcDeckID == "" {
		return errors.New("npc deck ID is required")
	}

	return nil
}
