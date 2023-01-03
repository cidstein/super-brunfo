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

type MatchRepositoryTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *MatchRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	conn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := pgx.Connect(suite.ctx, conn)
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

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

	m, err := matchRepo.FindByID(suite.ctx, match.ID)
	suite.NoError(err)
	suite.Equal(match.ID, m.ID)

	err = matchRepo.Delete(suite.ctx, match.ID)
	suite.NoError(err)
}
