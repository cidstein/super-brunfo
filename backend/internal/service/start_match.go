package service

import (
	"context"
	"errors"

	"github.com/cidstein/super-brunfo/internal/model"
	"github.com/cidstein/super-brunfo/internal/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type StartMatchUseCase struct {
	CardRepository  repository.CardRepositoryInterface
	DeckRepository  repository.DeckRepositoryInterface
	MatchRepository repository.MatchRepositoryInterface
	RoundRepository repository.RoundRepositoryInterface
}

/*
Start starts a new match
- Find all cards
- Split cards in two decks
- Shuffle decks
- Save decks
- Create match
- Create all the rounds without playing them
- Return match
*/
func (s *StartMatchUseCase) Start(ctx context.Context, db *pgx.Conn) (MatchOutputDTO, error) {
	s.CardRepository = repository.NewCardRepository(db)
	s.DeckRepository = repository.NewDeckRepository(db)
	s.MatchRepository = repository.NewMatchRepository(db)
	s.RoundRepository = repository.NewRoundRepository(db)

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

	err = s.MatchRepository.Save(ctx, match)
	if err != nil {
		msg := "Error saving match: " + err.Error()
		return MatchOutputDTO{}, errors.New(msg)
	}

	for i := 0; i < 10; i++ {
		roundID := uuid.New().String()
		/*
			Victory and attribute are not set here because they are set when the round is played
		*/
		round := model.NewRound(roundID, match.ID, playerDeck.Cards[i].ID, npcDeck.Cards[i].ID, i+1, false, "")

		err = s.RoundRepository.Save(ctx, round)
		if err != nil {
			msg := "Error saving round: " + err.Error()
			return MatchOutputDTO{}, errors.New(msg)
		}
	}

	return MatchOutputDTO{
		ID: match.ID,
	}, nil
}
