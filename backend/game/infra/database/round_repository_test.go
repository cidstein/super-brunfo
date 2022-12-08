package database

import (
	"context"

	"github.com/cidstein/super-brunfo/game/entity"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"
)

type RoundRepositoryTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *RoundRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	db, err := pgx.Connect(suite.ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	suite.NoError(err)
	suite.Db = db
}

func (suite *RoundRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func (suite *RoundRepositoryTestSuite) TestGivenAnRound_WhenSave_ThenShouldSaveRound() {
	roundID := uuid.New().String()
	matchID := uuid.New().String()

	cardRepo := NewCardRepository(suite.Db)
	cards, err := cardRepo.FindAll(suite.ctx)
	suite.NoError(err)

	deckRepo := NewDeckRepository(suite.Db)
	cut := len(cards) / 2
	pd := entity.Deck{
		ID:    uuid.New().String(),
		Cards: cards[:cut],
	}

	nd := entity.Deck{
		ID:    uuid.New().String(),
		Cards: cards[cut:],
	}

	playerDeck, err := deckRepo.Save(suite.ctx, pd)
	suite.NoError(err)

	npcDeck, err := deckRepo.Save(suite.ctx, nd)
	suite.NoError(err)

	match := entity.NewMatch(matchID, playerDeck.ID, npcDeck.ID, false, false)
	suite.NoError(match.IsValid())

	matchRepo := NewMatchRepository(suite.Db)
	err = matchRepo.Save(suite.ctx, match)
	suite.NoError(err)

	playerCard, err := playerDeck.Draw()
	suite.NoError(err)

	npcCard, err := npcDeck.Draw()
	suite.NoError(err)

	round := entity.NewRound(roundID, match.ID, playerCard.ID, npcCard.ID, false, "attack")
	suite.NoError(round.IsValid())

	roundRepo := NewRoundRepository(suite.Db)
	err = roundRepo.Save(suite.ctx, round)
	suite.NoError(err)

	victory, err := playerCard.Combat(&npcCard, "attack")
	suite.NoError(err)
	round.Victory = victory

	err = roundRepo.Update(suite.ctx, round)
	suite.NoError(err)
}
