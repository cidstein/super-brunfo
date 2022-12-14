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
