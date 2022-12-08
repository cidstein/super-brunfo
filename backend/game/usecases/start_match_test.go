package usecases

import (
	"context"

	"github.com/cidstein/super-brunfo/game/infra/database"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"
)

type StartMatchTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *StartMatchTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	db, err := pgx.Connect(suite.ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
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
