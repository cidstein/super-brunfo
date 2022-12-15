package model

import "errors"

type Match struct {
	ID           string `json:"id"`
	PlayerDeckID string `json:"playerDeckID"`
	NpcDeckID    string `json:"npcDeckID"`
	Victory      bool   `json:"victory"`
	Finished     bool   `json:"finished"`
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
