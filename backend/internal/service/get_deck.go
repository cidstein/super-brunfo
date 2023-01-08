package service

import (
	"context"

	"github.com/cidstein/super-brunfo/internal/repository"
)

type GetDeckUseCase struct {
	CardRepository repository.CardRepositoryInterface
	DeckRepository repository.DeckRepositoryInterface
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
			ImageURL:     card.ImageURL,
		}
	}

	return DeckOutputDTO{
		ID:    deck.ID,
		Cards: cards,
	}, nil
}
