package service

import (
	"context"
	"fmt"
	"os"

	"github.com/cidstein/super-brunfo/internal/repository"
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

func (suite *ListCardsTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func (suite *ListCardsTestSuite) TestGivenNothing_WhenListCards_ShouldReceiveListCards() {
	lcuc := ListCardsUseCase{}
	lcuc.CardRepository = repository.NewCardRepository(suite.Db)

	cards, err := lcuc.ListCards(suite.ctx, suite.Db)
	suite.NoError(err)
	suite.NotEmpty(cards)
}
