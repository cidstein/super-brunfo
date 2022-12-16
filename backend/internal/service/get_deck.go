package service

import (
	"context"

	"github.com/cidstein/super-brunfo/internal/infra/database"
)

type GetDeckUseCase struct {
	CardRepository database.CardRepositoryInterface
	DeckRepository database.DeckRepositoryInterface
}

func (s *GetDeckUseCase) GetDeck(ctx context.Context, deckID string) (DeckOutputDTO, error) {
	deck, err := s.DeckRepository.FindByID(ctx, deckID)
	if err != nil {
		return DeckOutputDTO{}, err
	}

	cards := make([]CardOutputDTO, len(deck.Cards))
	for i, card := range deck.Cards {
		cards[i] = CardOutputDTO{
			ID:           card.ID,
			Name:         card.Name,
			Attack:       card.Attack,
			Defense:      card.Defense,
			Intelligence: card.Intelligence,
			Agility:      card.Agility,
			Resilience:   card.Resilience,
		}
	}

	return DeckOutputDTO{
		ID:    deck.ID,
		Cards: cards,
	}, nil
}