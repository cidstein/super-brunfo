package database

import (
	"context"
	"fmt"
	"os"

	"github.com/cidstein/super-brunfo/game/entity"
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
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	suite.ctx = context.Background()
	db, err := pgx.Connect(suite.ctx, conn)
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
