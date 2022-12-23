package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/cidstein/super-brunfo/internal/infra/database"
	"github.com/jackc/pgx/v5"
)

type PlayRoundUseCase struct {
	CardRepository  database.CardRepositoryInterface
	RoundRepository database.RoundRepositoryInterface
}

func (p *PlayRoundUseCase) Play(ctx context.Context, db *pgx.Conn, roundID, attribute string) (RoundOutputDTO, error) {
	p.CardRepository = database.NewCardRepository(db)
	p.RoundRepository = database.NewRoundRepository(db)

	round, err := p.RoundRepository.FindByID(ctx, roundID)
	if err != nil {
		msg := fmt.Sprintf("error finding round: %v", err)
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

	round.Counter++
	round.Victory = victory
	round.Attribute = attribute

	err = p.RoundRepository.Update(ctx, *round)
	if err != nil {
		msg := fmt.Sprintf("error updating round: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	return RoundOutputDTO{
		ID: round.ID,
		PlayerCard: CardOutputDTO{
			ID: playerCard.ID,
		},
		NpcCard: CardOutputDTO{
			ID: npcCard.ID,
		},
		Counter:   round.Counter,
		Victory:   round.Victory,
		Attribute: round.Attribute,
	}, nil
}
