package model

import "errors"

type Round struct {
	ID           string `json:"id"`
	MatchID      string `json:"matchID"`
	PlayerCardID string `json:"playerCardID"`
	NpcCardID    string `json:"npcCardID"`
	Victory      bool   `json:"victory"`
	Attribute    string `json:"attribute"`
}

func NewRound(id string, matchID string, playerCardID string, npcCardID string, victory bool, attribute string) Round {
	return Round{
		ID:           id,
		MatchID:      matchID,
		PlayerCardID: playerCardID,
		NpcCardID:    npcCardID,
		Victory:      victory,
		Attribute:    attribute,
	}
}

func (r Round) IsValid() error {
	if r.ID == "" {
		return errors.New("ID is required")
	}

	if r.MatchID == "" {
		return errors.New("match ID is required")
	}

	if r.PlayerCardID == "" {
		return errors.New("player card ID is required")
	}

	if r.NpcCardID == "" {
		return errors.New("npc card ID is required")
	}

	if r.Attribute == "" {
		return errors.New("attribute is required")
	}

	return nil
}
