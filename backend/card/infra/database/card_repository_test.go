package database

import (
	"context"
	"testing"

	"github.com/cidstein/super-brunfo/card/entity"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"
)

type CardRepositoryTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *CardRepositoryTestSuite) SetupSuite() {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	suite.NoError(err)
	suite.Db = db
}

func (suite *CardRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CardRepositoryTestSuite))
}

func (suite *CardRepositoryTestSuite) TestGivenAnCard_WhenSave_ThenShouldSaveCard() {
	id := uuid.New().String()
	card := entity.NewCard(id, "name", 0, 0, 0, 0, 0)
	suite.NoError(card.IsValid())
	repo := NewCardRepository(suite.Db)
	err := repo.Save(card)
	suite.NoError(err)

	cardResult, err := repo.FindByID(card.ID)
	suite.NoError(err)
	suite.Equal(card.ID, cardResult.ID)
	suite.Equal(card.Name, cardResult.Name)
	suite.Equal(card.Attack, cardResult.Attack)
	suite.Equal(card.Defense, cardResult.Defense)
	suite.Equal(card.Intelligence, cardResult.Intelligence)
	suite.Equal(card.Agility, cardResult.Agility)
	suite.Equal(card.Resilience, cardResult.Resilience)
}
