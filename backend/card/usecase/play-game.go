package usecase

import (
	"github.com/cidstein/super-brunfo/card/entity"
	"github.com/google/uuid"
)

type PlayGameUseCase struct {
	DeckRepository  entity.DeckRepositoryInterface
	MatchRepository entity.MatchRepositoryInterface
	RoundRepository entity.RoundRepositoryInterface
}

func (p *PlayGameUseCase) Play(matchID string) (entity.Match, error) {
	/*
		1. Get match
		2. Check if decks are empty
		3. If decks are empty, compute winner
		4. If decks are not empty, create round
		5. Compute round
		6. Start over again

	*/

	match, err := p.MatchRepository.FindByID(matchID)
	if err != nil {
		return match, err
	}

	playerDeck, err := p.DeckRepository.FindByID(match.PlayerDeckID)
	if err != nil {
		return match, err
	}

	npcDeck, err := p.DeckRepository.FindByID(match.NpcDeckID)
	if err != nil {
		return match, err
	}

	for !playerDeck.CheckIfEmpty() {
		/*
			1. Create round
			2. Determine who chooses the attribute
			3. Receive attribute
			4. Compare attributes
			5. Determine winner
		*/

		round := entity.Round{
			ID:           uuid.New().String(),
			MatchID:      match.ID,
			PlayerCardID: playerDeck.Cards[0].ID,
			NpcCardID:    npcDeck.Cards[0].ID,
			Victory:      false,
			Attribute:    "attack",
		}

		err := p.RoundRepository.Save(round)
		if err != nil {
			return match, err
		}

	}

	match, err = p.MatchRepository.ComputeWinner(match)
	if err != nil {
		return match, err
	}

	return match, nil
}
