package database

import (
	"context"

	"github.com/cidstein/super-brunfo/card/entity"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"
)

type MatchRepositoryTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *MatchRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	db, err := pgx.Connect(suite.ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	suite.NoError(err)
	suite.Db = db
}

func (suite *MatchRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func (suite *MatchRepositoryTestSuite) TestGivenAnMatch_WhenSave_ThenShouldSaveMatch() {
	matchID := uuid.New().String()

	deckRepo := NewDeckRepository(suite.Db)

	deck, err := deckRepo.Save(suite.ctx)
	suite.NoError(err)

	playerDeck, npcDeck, err := deck.Split()
	suite.NoError(err)

	match := entity.NewMatch(matchID, playerDeck.ID, npcDeck.ID, false, false)
	suite.NoError(match.IsValid())
	repo := NewMatchRepository(suite.Db)

	err = repo.Save(suite.ctx, match)
	suite.NoError(err)

	m, err := repo.FindByID(suite.ctx, match.ID)
	suite.NoError(err)
	suite.Equal(match.ID, m.ID)

	err = repo.Delete(suite.ctx, deck.ID)
	suite.NoError(err)
}
