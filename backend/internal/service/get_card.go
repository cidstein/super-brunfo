package service

import (
	"context"

	"github.com/cidstein/super-brunfo/internal/infra/database"
	"github.com/jackc/pgx/v5"
)

type GetCardUseCase struct {
	CardRepository database.CardRepositoryInterface
}

func (s *GetCardUseCase) GetCard(ctx context.Context, db *pgx.Conn, cardID string) (CardOutputDTO, error) {
	s.CardRepository = database.NewCardRepository(db)
	card, err := s.CardRepository.FindByID(ctx, cardID)
	if err != nil {
		return CardOutputDTO{}, err
	}

	cardDTO := CardOutputDTO{
		ID:           card.ID,
		Name:         card.Name,
		Attack:       card.Attack,
		Defense:      card.Defense,
		Intelligence: card.Intelligence,
		Agility:      card.Agility,
		Resilience:   card.Resilience,
		ImageURL:     card.ImageURL,
	}

	return cardDTO, nil
}
