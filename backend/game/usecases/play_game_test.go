package usecases

import (
	"context"
	"fmt"
	"os"

	"github.com/cidstein/super-brunfo/game/entity"
	"github.com/cidstein/super-brunfo/game/infra/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type PlayGameTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *PlayGameTestSuite) SetupSuite() {
	err := godotenv.Load("../../.env")
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

func (suite *PlayGameTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func (suite *PlayGameTestSuite) TestGivenMatch_WhenCreateNewMatch_ShouldReceiveError() {
	pguc := PlayGameUseCase{}
	pguc.CardRepository = database.NewCardRepository(suite.Db)
	pguc.DeckRepository = database.NewDeckRepository(suite.Db)
	pguc.MatchRepository = database.NewMatchRepository(suite.Db)
	pguc.RoundRepository = database.NewRoundRepository(suite.Db)

	matchID := uuid.New().String()

	cards, err := pguc.CardRepository.FindAll(suite.ctx)
	suite.NoError(err)

	cut := len(cards) / 2
	pd := entity.Deck{
		ID:    uuid.New().String(),
		Cards: cards[:cut],
	}

	nd := entity.Deck{
		ID:    uuid.New().String(),
		Cards: cards[cut:],
	}

	playerDeck, err := pguc.DeckRepository.Save(suite.ctx, pd)
	suite.NoError(err)

	npcDeck, err := pguc.DeckRepository.Save(suite.ctx, nd)
	suite.NoError(err)

	match := entity.NewMatch(matchID, playerDeck.ID, npcDeck.ID, false, false)
	suite.NoError(match.IsValid())

	err = pguc.MatchRepository.Save(suite.ctx, match)
	suite.NoError(err)

	_, err = pguc.Play(suite.ctx, suite.Db, matchID, "attack")
	suite.NoError(err)
}
