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
