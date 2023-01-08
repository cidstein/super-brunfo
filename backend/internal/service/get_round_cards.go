package service

import (
	"context"

	"github.com/cidstein/super-brunfo/internal/repository"
	"github.com/jackc/pgx/v5"
)

type GetRoundCardsUseCase struct {
	RoundRepository repository.RoundRepositoryInterface
}

func (s *GetRoundCardsUseCase) GetRoundCards(ctx context.Context, db *pgx.Conn, roundID string) ([]CardOutputDTO, error) {
	s.RoundRepository = repository.NewRoundRepository(db)
	cards, err := s.RoundRepository.FindCardsByID(ctx, roundID)
	if err != nil {
		return nil, err
	}

	cardsDTO := make([]CardOutputDTO, len(cards))
	for i, card := range cards {
		cardsDTO[i] = CardOutputDTO{
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

	return cardsDTO, nil
}
