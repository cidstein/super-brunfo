package service

import (
	"context"
	"fmt"
	"os"

	"github.com/cidstein/super-brunfo/internal/model"
	"github.com/cidstein/super-brunfo/internal/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type LoadRoundTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *LoadRoundTestSuite) SetupSuite() {
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
		fmt.Println("Error connecting to repository")
		panic(err)
	}

	suite.Db = db
}

func (suite *LoadRoundTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func (suite *LoadRoundTestSuite) TestGivenMatch_WhenCreateNewMatch_ShouldReceiveError() {
	pguc := LoadRoundUseCase{}
	pguc.CardRepository = repository.NewCardRepository(suite.Db)
	pguc.DeckRepository = repository.NewDeckRepository(suite.Db)
	pguc.MatchRepository = repository.NewMatchRepository(suite.Db)
	pguc.RoundRepository = repository.NewRoundRepository(suite.Db)

	matchID := uuid.New().String()

	cards, err := pguc.CardRepository.FindAll(suite.ctx)
	suite.NoError(err)

	cut := len(cards) / 2
	pd := model.Deck{
		ID:    uuid.New().String(),
		Cards: cards[:cut],
	}

	nd := model.Deck{
		ID:    uuid.New().String(),
		Cards: cards[cut:],
	}

	playerDeck, err := pguc.DeckRepository.Save(suite.ctx, pd)
	suite.NoError(err)

	npcDeck, err := pguc.DeckRepository.Save(suite.ctx, nd)
	suite.NoError(err)

	match := model.NewMatch(matchID, playerDeck.ID, npcDeck.ID, false, false)
	suite.NoError(match.IsValid())

	err = pguc.MatchRepository.Save(suite.ctx, match)
	suite.NoError(err)

	_, err = pguc.LoadRound(suite.ctx, suite.Db, matchID)
	suite.NoError(err)
}
