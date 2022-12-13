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

type MatchRepositoryTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *MatchRepositoryTestSuite) SetupSuite() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	dbUser, ok := os.LookupEnv("DB_USER")
	suite.True(ok)

	dbPass, ok := os.LookupEnv("DB_PASSWORD")
	suite.True(ok)

	dbHost, ok := os.LookupEnv("DB_HOST")
	suite.True(ok)

	dbPort, ok := os.LookupEnv("DB_PORT")
	suite.True(ok)

	dbName, ok := os.LookupEnv("DB_NAME")
	suite.True(ok)

	conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	suite.ctx = context.Background()
	db, err := pgx.Connect(suite.ctx, conn)
	suite.NoError(err)
	suite.Db = db
}

func (suite *MatchRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func (suite *MatchRepositoryTestSuite) TestGivenAnMatch_WhenSave_ThenShouldSaveMatch() {
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

	m, err := matchRepo.FindByID(suite.ctx, match.ID)
	suite.NoError(err)
	suite.Equal(match.ID, m.ID)

	err = matchRepo.Delete(suite.ctx, match.ID)
	suite.NoError(err)
}
