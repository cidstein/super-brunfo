package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/cidstein/super-brunfo/internal/infra/database"
	"github.com/cidstein/super-brunfo/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PlayGameUseCase struct {
	CardRepository  database.CardRepositoryInterface
	DeckRepository  database.DeckRepositoryInterface
	MatchRepository database.MatchRepositoryInterface
	RoundRepository database.RoundRepositoryInterface
}

func (p *PlayGameUseCase) Play(ctx context.Context, db *pgx.Conn, matchID, attribute string) (RoundOutputDTO, error) {
	p.MatchRepository = database.NewMatchRepository(db)
	match, err := p.MatchRepository.FindByID(ctx, matchID)
	if err != nil {
		msg := fmt.Sprintf("error finding match: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	if match.Finished {
		msg := fmt.Sprintf("match %s is already finished", matchID)
		return RoundOutputDTO{}, errors.New(msg)
	}

	p.DeckRepository = database.NewDeckRepository(db)
	playerDeck, err := p.DeckRepository.FindByID(ctx, match.PlayerDeckID)
	if err != nil {
		msg := fmt.Sprintf("error finding player deck: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	if playerDeck.CheckIfEmpty() {
		match, err = p.MatchRepository.ComputeWinner(ctx, *match)
		if err != nil {
			msg := fmt.Sprintf("error computing winner: %v", err)
			return RoundOutputDTO{}, errors.New(msg)
		}
	}

	npcDeck, err := p.DeckRepository.FindByID(ctx, match.NpcDeckID)
	if err != nil {
		msg := fmt.Sprintf("error finding npc deck: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	playerCardID := playerDeck.Cards[0].ID
	npcCardID := npcDeck.Cards[0].ID

	round := model.Round{
		ID:           uuid.New().String(),
		MatchID:      match.ID,
		PlayerCardID: playerCardID,
		NpcCardID:    npcCardID,
		Victory:      false,
		Attribute:    attribute,
	}

	p.RoundRepository = database.NewRoundRepository(db)
	err = p.RoundRepository.Save(ctx, round)
	if err != nil {
		msg := fmt.Sprintf("error saving round: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	p.CardRepository = database.NewCardRepository(db)
	playerCard, err := p.CardRepository.FindByID(ctx, playerCardID)
	if err != nil {
		msg := fmt.Sprintf("error finding player card: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	npcCard, err := p.CardRepository.FindByID(ctx, npcCardID)
	if err != nil {
		msg := fmt.Sprintf("error finding npc card: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	roundWon, err := playerCard.Combat(npcCard, attribute)
	if err != nil {
		msg := fmt.Sprintf("error comparing cards: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	round.Victory = roundWon

	err = p.RoundRepository.Update(ctx, round)
	if err != nil {
		msg := fmt.Sprintf("error updating round: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	err = p.DeckRepository.DrawCard(ctx, match.PlayerDeckID, playerCard.ID)
	if err != nil {
		msg := fmt.Sprintf("error drawing card from player deck: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	err = p.DeckRepository.DrawCard(ctx, match.NpcDeckID, npcCard.ID)
	if err != nil {
		msg := fmt.Sprintf("error drawing card from npc deck: %v", err)
		return RoundOutputDTO{}, errors.New(msg)
	}

	return RoundOutputDTO{
		ID: round.ID,
		Match: MatchOutputDTO{
			ID: match.ID,
			PlayerDeck: DeckOutputDTO{
				ID: match.PlayerDeckID,
			},
			NpcDeck: DeckOutputDTO{
				ID: match.NpcDeckID,
			},
			Finished: match.Finished,
			Victory:  match.Victory,
		},
		PlayerCard: CardOutputDTO{
			ID:           playerCard.ID,
			Name:         playerCard.Name,
			Attack:       playerCard.Attack,
			Defense:      playerCard.Defense,
			Intelligence: playerCard.Intelligence,
			Agility:      playerCard.Agility,
			Resilience:   playerCard.Resilience,
		},
		NpcCard: CardOutputDTO{
			ID:           npcCard.ID,
			Name:         npcCard.Name,
			Attack:       npcCard.Attack,
			Defense:      npcCard.Defense,
			Intelligence: npcCard.Intelligence,
			Agility:      npcCard.Agility,
			Resilience:   npcCard.Resilience,
		},
		Victory:   roundWon,
		Attribute: attribute,
	}, nil
}
