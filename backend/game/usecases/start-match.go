package usecases

import (
	"context"

	"github.com/cidstein/super-brunfo/game/entity"
	"github.com/cidstein/super-brunfo/game/infra/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
	CardRepository  entity.CardRepositoryInterface
	DeckRepository  entity.DeckRepositoryInterface
	MatchRepository entity.MatchRepositoryInterface
}

func (s *StartMatchUseCase) Start(ctx context.Context, db *pgx.Conn) (MatchOutputDTO, error) {
	s.CardRepository = database.NewCardRepository(db)
	cards, err := s.CardRepository.FindAll(ctx)
	if err != nil {
		return MatchOutputDTO{}, err
	}

	cut := len(cards) / 2

	s.DeckRepository = database.NewDeckRepository(db)
	playerDeck, err := s.DeckRepository.Save(ctx, cards[:cut])
	if err != nil {
		return MatchOutputDTO{}, err
	}

	npcDeck, err := s.DeckRepository.Save(ctx, cards[cut:])
	if err != nil {
		return MatchOutputDTO{}, err
	}

	playerDeck.Shuffle()
	npcDeck.Shuffle()

	id := uuid.New().String()
	match := entity.NewMatch(id, playerDeck.ID, npcDeck.ID, false, false)

	s.MatchRepository = database.NewMatchRepository(db)

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
			ID: npcDeck.ID,
		},
		Winner: false,
	}, nil
}
