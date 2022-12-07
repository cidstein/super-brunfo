package usecases

import (
	"context"

	"github.com/cidstein/super-brunfo/game/entity"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PlayGameUseCase struct {
	CardRepository  entity.CardRepositoryInterface
	DeckRepository  entity.DeckRepositoryInterface
	MatchRepository entity.MatchRepositoryInterface
	RoundRepository entity.RoundRepositoryInterface
}

func (p *PlayGameUseCase) Play(ctx context.Context, db *pgx.Conn, matchID, attribute string) (entity.Match, error) {
	/*
		1. Get match
		2. Check if decks are empty
		3. If decks are empty, compute winner
		4. If decks are not empty, create round
		5. Compute round
		6. Start over again

	*/

	match, err := p.MatchRepository.FindByID(ctx, matchID)
	if err != nil {
		return match, err
	}

	if match.Finished {
		return match, nil
	}

	playerDeck, err := p.DeckRepository.FindByID(ctx, match.PlayerDeckID)
	if err != nil {
		return match, err
	}

	if playerDeck.CheckIfEmpty() {
		match, err = p.MatchRepository.ComputeWinner(ctx, match)
		if err != nil {
			return match, err
		}
	}

	npcDeck, err := p.DeckRepository.FindByID(ctx, match.NpcDeckID)
	if err != nil {
		return match, err
	}

	/*
		1. Create round
		2. Determine who chooses the attribute
		3. Receive attribute
		4. Compare attributes
		5. Determine winner
	*/

	playerCardID := playerDeck.Cards[0].ID
	npcCardID := npcDeck.Cards[0].ID

	round := entity.Round{
		ID:           uuid.New().String(),
		MatchID:      match.ID,
		PlayerCardID: playerCardID,
		NpcCardID:    npcCardID,
		Victory:      false,
		Attribute:    attribute,
	}

	err = p.RoundRepository.Save(ctx, round)
	if err != nil {
		return match, err
	}

	playerCard, err := p.CardRepository.FindByID(ctx, playerCardID)
	if err != nil {
		return match, err
	}

	npcCard, err := p.CardRepository.FindByID(ctx, npcCardID)
	if err != nil {
		return match, err
	}

	roundWon, err := playerCard.Combat(npcCard, attribute)
	if err != nil {
		return match, err
	}

	round.Victory = roundWon

	err = p.RoundRepository.Update(ctx, round)
	if err != nil {
		return match, err
	}

	return match, nil
}
