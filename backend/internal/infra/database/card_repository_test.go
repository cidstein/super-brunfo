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

type CardRepositoryTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *CardRepositoryTestSuite) SetupSuite() {
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

func (suite *CardRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func (suite *CardRepositoryTestSuite) TestGivenAnCard_WhenSave_ThenShouldSaveCard() {
	id := uuid.New().String()
	card := model.NewCard(id, "name", 0, 0, 0, 0, 0, "https://1drv.ms/u/s!Aq4ssY1EMmRWh9o6Ax8aMkel7eetng?e=s1TMI2")
	suite.NoError(card.IsValid())
	repo := NewCardRepository(suite.Db)

	err := repo.Save(suite.ctx, card)
	suite.NoError(err)

	cardResult, err := repo.FindByID(suite.ctx, card.ID)
	suite.NoError(err)
	suite.Equal(card.ID, cardResult.ID)
	suite.Equal(card.Name, cardResult.Name)
	suite.Equal(card.Attack, cardResult.Attack)
	suite.Equal(card.Defense, cardResult.Defense)
	suite.Equal(card.Intelligence, cardResult.Intelligence)
	suite.Equal(card.Agility, cardResult.Agility)
	suite.Equal(card.Resilience, cardResult.Resilience)
	suite.Equal(card.ImageURL, cardResult.ImageURL)

	err = repo.Delete(suite.ctx, card.ID)
	suite.NoError(err)
}
