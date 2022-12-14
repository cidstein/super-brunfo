package database

import (
	"context"
	"fmt"
	"os"

	"github.com/cidstein/super-brunfo/game/entity"
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
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	// dbUser, ok := os.LookupEnv("DB_USER")
	// suite.True(ok)

	// dbPass, ok := os.LookupEnv("DB_PASSWORD")
	// suite.True(ok)

	// dbHost, ok := os.LookupEnv("DB_HOST")
	// suite.True(ok)

	// dbPort, ok := os.LookupEnv("DB_PORT")
	// suite.True(ok)

	// dbName, ok := os.LookupEnv("DB_NAME")
	// suite.True(ok)

	dbUrl, ok := os.LookupEnv("DB_URL")
	suite.True(ok)

	// conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	// 	dbUser,
	// 	dbPass,
	// 	dbHost,
	// 	dbPort,
	// 	dbName,
	// )

	suite.ctx = context.Background()
	db, err := pgx.Connect(suite.ctx, dbUrl)
	suite.NoError(err)
	suite.Db = db
}

func (suite *CardRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func (suite *CardRepositoryTestSuite) TestGivenAnCard_WhenSave_ThenShouldSaveCard() {
	id := uuid.New().String()
	card := entity.NewCard(id, "name", 0, 0, 0, 0, 0)
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

	err = repo.Delete(suite.ctx, card.ID)
	suite.NoError(err)
}
