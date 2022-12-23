package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/cidstein/super-brunfo/internal/infra/database"
	"github.com/jackc/pgx/v5"
)

type LoadRoundUseCase struct {
	CardRepository  database.CardRepositoryInterface
	DeckRepository  database.DeckRepositoryInterface
	MatchRepository database.MatchRepositoryInterface
	RoundRepository database.RoundRepositoryInterface
}

/*
Play loads a round
- Find match
- Check if match is finished
- Find the round to be played
- Find player card
- Find npc card
- Return round
*/
func (p *LoadRoundUseCase) LoadRound(ctx context.Context, db *pgx.Conn, matchID string) (RoundOutputDTO, error) {
	p.CardRepository = database.NewCardRepository(db)
	p.DeckRepository = database.NewDeckRepository(db)
	p.MatchRepository = database.NewMatchRepository(db)
	p.RoundRepository = database.NewRoundRepository(db)

	match, err := p.MatchRepository.FindByID(ctx, matchID)
	if err != nil {
		msg := fmt.Sprintf("error finding match: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	if match.Finished {
		msg := fmt.Sprintf("match %s is already finished", matchID)
		return RoundOutputDTO{}, errors.New(msg)
	}

	round, err := p.RoundRepository.FindRoundToBePlayed(ctx, matchID)
	if err != nil {
		msg := fmt.Sprintf("error finding round to be played: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	playerCard, err := p.CardRepository.FindByID(ctx, round.PlayerCardID)
	if err != nil {
		msg := fmt.Sprintf("error finding player card: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	npcCard, err := p.CardRepository.FindByID(ctx, round.NpcCardID)
	if err != nil {
		msg := fmt.Sprintf("error finding npc card: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	return RoundOutputDTO{
		ID: round.ID,
		Match: MatchOutputDTO{
			ID:       match.ID,
			Counter:  match.Counter,
			Victory:  match.Victory,
			Finished: match.Finished,
		},
		Cards: []CardOutputDTO{
			{
				ID:           playerCard.ID,
				Name:         playerCard.Name,
				Attack:       playerCard.Attack,
				Defense:      playerCard.Defense,
				Intelligence: playerCard.Intelligence,
				Agility:      playerCard.Agility,
				Resilience:   playerCard.Resilience,
				ImageURL:     playerCard.ImageURL,
			},
			{
				ID:           npcCard.ID,
				Name:         npcCard.Name,
				Attack:       npcCard.Attack,
				Defense:      npcCard.Defense,
				Intelligence: npcCard.Intelligence,
				Agility:      npcCard.Agility,
				Resilience:   npcCard.Resilience,
				ImageURL:     npcCard.ImageURL,
			},
		},
		Counter:   round.Counter,
		Victory:   round.Victory,
		Attribute: round.Attribute,
	}, nil
}
