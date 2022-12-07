package usecase

import (
	"context"

	"github.com/cidstein/super-brunfo/game/entity"
	"github.com/google/uuid"
)

type CardOutputDTO struct {
	ID           string
	Name         string
	Attack       int
	Defense      int
	Intelligence int
	Agility      int
	Resilience   int
}

type DeckOutputDTO struct {
	ID    string
	Cards []CardOutputDTO
}

type MatchOutputDTO struct {
	ID         string
	PlayerDeck DeckOutputDTO
	ComDeck    DeckOutputDTO
	Winner     bool
}

type StartMatchUseCase struct {
	DeckRepository  entity.DeckRepositoryInterface
	MatchRepository entity.MatchRepositoryInterface
}

func (s *StartMatchUseCase) Start(ctx context.Context) (MatchOutputDTO, error) {
	deck, err := s.DeckRepository.Save(ctx)
	if err != nil {
		return MatchOutputDTO{}, err
	}

	deck.Shuffle()
	playerDeck, comDeck, err := deck.Split()
	if err != nil {
		return MatchOutputDTO{}, err
	}

	id := uuid.New().String()
	match := entity.NewMatch(id, playerDeck.ID, comDeck.ID, false, false)

	err = s.MatchRepository.Save(ctx, match)
	if err != nil {
		return MatchOutputDTO{}, err
	}

	return MatchOutputDTO{
		ID: match.ID,
		PlayerDeck: DeckOutputDTO{
			ID: playerDeck.ID,
		},
		ComDeck: DeckOutputDTO{
			ID: comDeck.ID,
		},
		Winner: false,
	}, nil
}
