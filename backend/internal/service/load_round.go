package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/cidstein/super-brunfo/internal/repository"
	"github.com/jackc/pgx/v5"
)

type LoadRoundUseCase struct {
	CardRepository  repository.CardRepositoryInterface
	DeckRepository  repository.DeckRepositoryInterface
	MatchRepository repository.MatchRepositoryInterface
	RoundRepository repository.RoundRepositoryInterface
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
	p.CardRepository = repository.NewCardRepository(db)
	p.DeckRepository = repository.NewDeckRepository(db)
	p.MatchRepository = repository.NewMatchRepository(db)
	p.RoundRepository = repository.NewRoundRepository(db)

	match, err := p.MatchRepository.FindByID(ctx, matchID)
	if err != nil {
		log.Printf("error finding match: %v", err)
		msg := fmt.Sprintf("error finding match: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	if match.Finished {
		log.Printf("match %s is already finished", matchID)
		msg := fmt.Sprintf("match %s is already finished", matchID)
		return RoundOutputDTO{}, errors.New(msg)
	}

	round, err := p.RoundRepository.FindRoundToBePlayed(ctx, matchID)
	if err != nil {
		log.Printf("error finding round to be played: %v", err)
		msg := fmt.Sprintf("error finding round to be played: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	playerCard, err := p.CardRepository.FindByID(ctx, round.PlayerCardID)
	if err != nil {
		log.Printf("error finding player card: %v", err)
		msg := fmt.Sprintf("error finding player card: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	npcCard, err := p.CardRepository.FindByID(ctx, round.NpcCardID)
	if err != nil {
		log.Printf("error finding npc card: %v", err)
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
				FlavourText:  playerCard.FlavourText,
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
				FlavourText:  npcCard.FlavourText,
				ImageURL:     npcCard.ImageURL,
			},
		},
		Counter:   round.Counter,
		Victory:   round.Victory,
		Attribute: round.Attribute,
	}, nil
}
