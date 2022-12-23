package database

import (
	"context"
	"fmt"
	"os"

	"github.com/cidstein/super-brunfo/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type RoundRepositoryTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *RoundRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	dbUrl := os.Getenv("DB_URL")
	suite.NotEmpty(dbUrl)

	db, err := pgx.Connect(suite.ctx, dbUrl)
	suite.NoError(err)
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

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
	pd := model.Deck{
		ID:    uuid.New().String(),
		Cards: cards[:cut],
	}

	nd := model.Deck{
		ID:    uuid.New().String(),
		Cards: cards[cut:],
	}

	playerDeck, err := deckRepo.Save(suite.ctx, pd)
	suite.NoError(err)

	npcDeck, err := deckRepo.Save(suite.ctx, nd)
	suite.NoError(err)

	match := model.NewMatch(matchID, playerDeck.ID, npcDeck.ID, false, false)
	suite.NoError(match.IsValid())

	matchRepo := NewMatchRepository(suite.Db)
	err = matchRepo.Save(suite.ctx, match)
	suite.NoError(err)

	playerCard, err := playerDeck.Draw()
	suite.NoError(err)

	npcCard, err := npcDeck.Draw()
	suite.NoError(err)

	round := model.NewRound(roundID, match.ID, playerCard.ID, npcCard.ID, 1, false, "attack")
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
