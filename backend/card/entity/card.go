package entity

import "errors"

type Card struct {
	ID           string
	Name         string
	Attack       int
	Defense      int
	Intelligence int
	Agility      int
	Resilience   int
}

func NewCard(id string, name string, attack int, defense int, intelligence int, agility int, resilience int) Card {
	return Card{
		ID:           id,
		Name:         name,
		Attack:       attack,
		Defense:      defense,
		Intelligence: intelligence,
		Agility:      agility,
		Resilience:   resilience,
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

	if c.Agility < 0 || c.Agility > 10 {
		return errors.New("agility must be between 0 and 100")
	}

	if c.Resilience < 0 || c.Resilience > 100 {
		return errors.New("resilience must be between 0 and 100")
	}

	return nil
}
