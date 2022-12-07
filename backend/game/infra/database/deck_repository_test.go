package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"
)

type DeckRepositoryTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *DeckRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	db, err := pgx.Connect(suite.ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	suite.NoError(err)
	suite.Db = db
}

func (suite *DeckRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func (suite *DeckRepositoryTestSuite) TestGivenAnDeck_WhenSave_ThenShouldSaveDeck() {
	cardRepo := NewCardRepository(suite.Db)
	cards, err := cardRepo.FindAll(suite.ctx)
	suite.NoError(err)

	repo := NewDeckRepository(suite.Db)

	deck, err := repo.Save(suite.ctx, cards)
	suite.NoError(err)

	d, err := repo.FindByID(suite.ctx, deck.ID)
	suite.NoError(err)
	suite.Equal(deck.ID, d.ID)

	err = repo.Delete(suite.ctx, deck.ID)
	suite.NoError(err)
}
