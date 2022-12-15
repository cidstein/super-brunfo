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

type GetDeckTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *GetDeckTestSuite) SetupSuite() {
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

func (suite *GetDeckTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func (suite *GetDeckTestSuite) TestGivenDeck_WhenGetDeck_ShouldReceiveDeck() {
	gduc := GetDeckUseCase{}
	gduc.DeckRepository = database.NewDeckRepository(suite.Db)
	gduc.CardRepository = database.NewCardRepository(suite.Db)

	cards, err := gduc.CardRepository.FindAll(suite.ctx)
	suite.NoError(err)

	cut := len(cards) / 2
	pd := entity.Deck{
		ID:    uuid.New().String(),
		Cards: cards[:cut],
	}

	d, err := gduc.DeckRepository.Save(suite.ctx, pd)
	suite.NoError(err)

	deck, err := gduc.GetDeck(suite.ctx, d.ID)
	suite.NoError(err)
	suite.Equal(deck.ID, d.ID)
	suite.Equal(deck.Cards, d.Cards)
}
