package service

import (
	"context"
	"fmt"
	"os"

	"github.com/cidstein/super-brunfo/internal/infra/database"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type ListCardsTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *ListCardsTestSuite) SetupSuite() {
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

func (suite *ListCardsTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func (suite *ListCardsTestSuite) TestGivenNothing_WhenListCards_ShouldReceiveListCards() {
	lcuc := ListCardsUseCase{}
	lcuc.CardRepository = database.NewCardRepository(suite.Db)

	cards, err := lcuc.ListCards(suite.ctx, suite.Db)
	suite.NoError(err)
	suite.NotEmpty(cards)
}
