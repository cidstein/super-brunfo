package database

import (
	"context"
	"fmt"
	"os"

	"github.com/cidstein/super-brunfo/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type DeckRepositoryTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *DeckRepositoryTestSuite) SetupSuite() {
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
		fmt.Println("Error connecting to database")
		panic(err)
	}

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
	deckDto := model.Deck{
		ID:    uuid.New().String(),
		Cards: cards,
	}

	deck, err := repo.Save(suite.ctx, deckDto)
	suite.NoError(err)

	d, err := repo.FindByID(suite.ctx, deck.ID)
	suite.NoError(err)
	suite.Equal(deck.ID, d.ID)

	err = repo.DrawCard(suite.ctx, deck.ID, cards[0].ID)
	suite.NoError(err)

	err = repo.Delete(suite.ctx, d.ID)
	suite.NoError(err)
}
