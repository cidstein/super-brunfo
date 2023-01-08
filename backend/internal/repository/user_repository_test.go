package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	ctx context.Context
	suite.Suite
	Db *pgx.Conn
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
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

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close(suite.ctx)
}

func (suite *UserRepositoryTestSuite) TestGivenAnUser_WhenSave_ThenShouldSaveUser() {
	repo := NewUserRepository(suite.Db)

	user, err := repo.FindByEmail(suite.ctx, "a@b.c")
	suite.NoError(err)
	suite.NotNil(user)
}
