package model

import (
	"errors"
)

type Card struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Attack       int    `json:"attack"`
	Defense      int    `json:"defense"`
	Intelligence int    `json:"intelligence"`
	Agility      int    `json:"agility"`
	Resilience   int    `json:"resilience"`
	ImageURL     string `json:"image_url"`
}

func NewCard(id string, name string, attack int, defense int, intelligence int, agility int, resilience int, imageUrl string) Card {
	return Card{
		ID:           id,
		Name:         name,
		Attack:       attack,
		Defense:      defense,
		Intelligence: intelligence,
		Agility:      agility,
		Resilience:   resilience,
		ImageURL:     imageUrl,
	}
}

func (c *Card) IsValid() error {
	if c.ID == "" {
		return errors.New("ID is required")
	}

	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.Attack < 0 || c.Attack > 100 {
		return errors.New("attack must be between 0 and 100")
	}

	if c.Defense < 0 || c.Defense > 100 {
		return errors.New("defense must be between 0 and 100")
	}

	if c.Intelligence < 0 || c.Intelligence > 100 {
		return errors.New("intelligence must be between 0 and 100")
	}

	if c.Agility < 0 || c.Agility > 100 {
		return errors.New("agility must be between 0 and 100")
	}

	if c.Resilience < 0 || c.Resilience > 100 {
		return errors.New("resilience must be between 0 and 100")
	}

	if c.ImageURL == "" {
		return errors.New("image_url is required")
	}

	return nil
}

func (c *Card) Combat(other *Card, attribute string) (bool, error) {
	var winner bool

	if c.IsValid() != nil {
		return winner, errors.New("player invalid card")
	}

	if other.IsValid() != nil {
		return winner, errors.New("npc invalid card")
	}

	switch attribute {
	case "attack":
		if c.Attack > other.Attack {
			winner = true
		}
	case "defense":
		if c.Defense > other.Defense {
			winner = true
		}
	case "intelligence":
		if c.Intelligence > other.Intelligence {
			winner = true
		}
	case "agility":
		if c.Agility > other.Agility {
			winner = true
		}
	case "resilience":
		if c.Resilience > other.Resilience {
			winner = true
		}
	default:
		return false, errors.New("invalid attribute")
	}

	return winner, nil
}
