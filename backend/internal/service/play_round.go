package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/cidstein/super-brunfo/internal/repository"
	"github.com/jackc/pgx/v5"
)

type PlayRoundUseCase struct {
	CardRepository  repository.CardRepositoryInterface
	MatchRepository repository.MatchRepositoryInterface
	RoundRepository repository.RoundRepositoryInterface
}

func (p *PlayRoundUseCase) Play(ctx context.Context, db *pgx.Conn, roundID, attribute string) (RoundOutputDTO, error) {
	p.CardRepository = repository.NewCardRepository(db)
	p.MatchRepository = repository.NewMatchRepository(db)
	p.RoundRepository = repository.NewRoundRepository(db)

	round, err := p.RoundRepository.FindByID(ctx, roundID)
	if err != nil {
		msg := fmt.Sprintf("error finding round: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	match, err := p.MatchRepository.FindByID(ctx, round.MatchID)
	if err != nil {
		msg := fmt.Sprintf("error finding match: %v", err)
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

	victory, err := playerCard.Combat(npcCard, attribute)
	if err != nil {
		msg := fmt.Sprintf("error playing round: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	round.Victory = victory
	round.Attribute = attribute

	err = p.RoundRepository.Update(ctx, *round)
	if err != nil {
		msg := fmt.Sprintf("error updating round: %v", err)
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
