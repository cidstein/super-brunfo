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

type ListMatchesTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *ListMatchesTestSuite) SetupSuite() {
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

func (suite *ListMatchesTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func (suite *ListMatchesTestSuite) TestGivenNothing_WhenListMatches_ShouldReceiveListMatches() {
	lmuc := ListMatchesUseCase{}
	lmuc.MatchRepository = database.NewMatchRepository(suite.Db)

	_, err := lmuc.ListMatches(suite.ctx)
	suite.NoError(err)
}
