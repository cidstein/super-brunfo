package service

import (
	"context"
	"errors"

	"github.com/cidstein/super-brunfo/internal/infra/database"
	"github.com/cidstein/super-brunfo/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

type StartMatchUseCase struct {
	CardRepository  database.CardRepositoryInterface
	DeckRepository  database.DeckRepositoryInterface
	MatchRepository database.MatchRepositoryInterface
}

func cardsDTO(cards []model.Card) []CardOutputDTO {
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
		msg := "Error finding cards: " + err.Error()
		return MatchOutputDTO{}, errors.New(msg)
	}

	cut := len(cards) / 2

	pdID := uuid.New().String()
	pd := model.NewDeck(pdID, cards[:cut])

	ndID := uuid.New().String()
	nd := model.NewDeck(ndID, cards[cut:])

	pd.Shuffle()
	nd.Shuffle()

	log.Info().Msgf("Player deck: %s, %v", pd.ID, pd.Cards)
	log.Info().Msgf("NPC deck: %s, %v", pd.ID, nd.Cards)

	cardsPlayer := cardsDTO(pd.Cards)
	cardsNpc := cardsDTO(nd.Cards)

	s.DeckRepository = database.NewDeckRepository(db)
	playerDeck, err := s.DeckRepository.Save(ctx, pd)
	if err != nil {
		msg := "Error saving player deck: " + err.Error()
		return MatchOutputDTO{}, errors.New(msg)
	}

	npcDeck, err := s.DeckRepository.Save(ctx, nd)
	if err != nil {
		msg := "Error saving npc deck: " + err.Error()
		return MatchOutputDTO{}, errors.New(msg)
	}

	id := uuid.New().String()
	match := model.NewMatch(id, playerDeck.ID, npcDeck.ID, false, false)

	s.MatchRepository = database.NewMatchRepository(db)

	err = s.MatchRepository.Save(ctx, match)
	if err != nil {
		msg := "Error saving match: " + err.Error()
		return MatchOutputDTO{}, errors.New(msg)
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
	}, nil
}
