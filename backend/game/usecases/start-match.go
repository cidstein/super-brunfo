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
	NpcDeck    DeckOutputDTO
	Winner     bool
}

type StartMatchUseCase struct {
	CardRepository  entity.CardRepositoryInterface
	DeckRepository  entity.DeckRepositoryInterface
	MatchRepository entity.MatchRepositoryInterface
}

func cardsDTO(cards []entity.Card) []CardOutputDTO {
	var cardsDTO []CardOutputDTO
	for card := range cards {
		cardsDTO = append(cardsDTO, CardOutputDTO{
			ID:           cards[card].ID,
			Name:         cards[card].Name,
			Attack:       cards[card].Attack,
			Defense:      cards[card].Defense,
			Intelligence: cards[card].Intelligence,
			Agility:      cards[card].Agility,
			Resilience:   cards[card].Resilience,
		})
	}
	return cardsDTO
}

func (s *StartMatchUseCase) Start(ctx context.Context, db *pgx.Conn) (MatchOutputDTO, error) {
	s.CardRepository = database.NewCardRepository(db)
	cards, err := s.CardRepository.FindAll(ctx)
	if err != nil {
		return MatchOutputDTO{}, err
	}

	cut := len(cards) / 2

	cardsPlayer := cardsDTO(cards[:cut])
	cardsNpc := cardsDTO(cards[cut:])

	s.DeckRepository = database.NewDeckRepository(db)
	playerDeck, err := s.DeckRepository.Save(ctx, cards[:cut])
	if err != nil {
		return MatchOutputDTO{}, err
	}

	npcDeck, err := s.DeckRepository.Save(ctx, cards[cut:])
	if err != nil {
		return MatchOutputDTO{}, err
	}

	// Not working yet, but know how to do it
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
			ID:    playerDeck.ID,
			Cards: cardsPlayer,
		},
		NpcDeck: DeckOutputDTO{
			ID:    npcDeck.ID,
			Cards: cardsNpc,
		},
		Winner: false,
	}, nil
}
