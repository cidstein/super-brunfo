package usecases

import (
	"context"
	"fmt"
	"os"

	"github.com/cidstein/super-brunfo/game/infra/database"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type StartMatchTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *StartMatchTestSuite) SetupSuite() {
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

func (suite *StartMatchTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func (suite *StartMatchTestSuite) TestGivenMatch_WhenCreateNewMatch_ShouldReceiveError() {
	smuc := StartMatchUseCase{}
	smuc.CardRepository = database.NewCardRepository(suite.Db)
	smuc.DeckRepository = database.NewDeckRepository(suite.Db)
	smuc.MatchRepository = database.NewMatchRepository(suite.Db)

	_, err := smuc.Start(suite.ctx, suite.Db)
	suite.NoError(err)
}
