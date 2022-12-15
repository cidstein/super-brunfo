package service

import (
	"context"

	"github.com/cidstein/super-brunfo/internal/infra/database"
	"github.com/jackc/pgx/v5"
)

type ListCardsUseCase struct {
	CardRepository database.CardRepositoryInterface
}

func (s *ListCardsUseCase) ListCards(ctx context.Context, db *pgx.Conn) ([]CardOutputDTO, error) {
	s.CardRepository = database.NewCardRepository(db)
	cards, err := s.CardRepository.FindAll(ctx)
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
		}
	}

	return cardsDTO, nil
}
